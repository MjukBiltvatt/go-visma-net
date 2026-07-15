package vismanet

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"sync/atomic"
	"testing"
	"time"
)

func TestGetCustomerInvoiceV1URL(t *testing.T) {
	c := NewClient(nil)

	// Do() defaults empty path params before building the URL, so the {{if .invoice_number}}
	// segment resolves. These cases call url() directly to stay offline, so they mimic that guard.
	buildURL := func(req GetCustomerInvoiceV1Request) (string, error) {
		if (*Request)(&req).pathParams == nil {
			req.SetPathParams(GetCustomerInvoiceV1PathParams{})
		}
		return (*Request)(&req).url()
	}

	// With query parameters set and no invoice number, the request targets the list/filter endpoint.
	list := c.NewGetCustomerInvoiceV1Request()
	list.SetQueryParams(GetCustomerInvoiceV1QueryParams{DocumentType: "Invoice", Status: "Open"})
	got, err := buildURL(list)
	if err != nil {
		t.Fatal(err)
	}
	if want := "https://api.finance.visma.net/v1/customerinvoice?documentType=Invoice&status=Open"; got != want {
		t.Errorf("list URL: expected %q, got %q", want, got)
	}

	// With an invoice number set, the request targets a single invoice and sends no query string.
	single := c.NewGetCustomerInvoiceV1Request()
	single.SetPathParams(GetCustomerInvoiceV1PathParams{InvoiceNumber: "12345"})
	got, err = buildURL(single)
	if err != nil {
		t.Fatal(err)
	}
	if want := "https://api.finance.visma.net/v1/customerinvoice/12345"; got != want {
		t.Errorf("single URL: expected %q, got %q", want, got)
	}

	// Empty query fields must be omitted entirely.
	partial := c.NewGetCustomerInvoiceV1Request()
	partial.SetQueryParams(GetCustomerInvoiceV1QueryParams{Status: "Open"})
	got, err = buildURL(partial)
	if err != nil {
		t.Fatal(err)
	}
	if want := "https://api.finance.visma.net/v1/customerinvoice?status=Open"; got != want {
		t.Errorf("partial URL: expected %q, got %q", want, got)
	}

	// Integer pagination params are rendered, and zero values omitted.
	paged := c.NewGetCustomerInvoiceV1Request()
	paged.SetQueryParams(GetCustomerInvoiceV1QueryParams{Status: "Open", PageNumber: 2, PageSize: 1000})
	got, err = buildURL(paged)
	if err != nil {
		t.Fatal(err)
	}
	if want := "https://api.finance.visma.net/v1/customerinvoice?pageNumber=2&pageSize=1000&status=Open"; got != want {
		t.Errorf("paged URL: expected %q, got %q", want, got)
	}
}

func TestGetCustomerInvoiceV1DoAll(t *testing.T) {
	const totalCount, maxPageSize = 7, 3

	// Serves the customerinvoice list endpoint with pagination and metadata, so DoAll can be
	// exercised without hitting the live API. The API's default page size equals maxPageSize.
	var requests int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		atomic.AddInt32(&requests, 1)
		q := req.URL.Query()

		// The caller's filters must be preserved on every page.
		if q.Get("documentType") != "Invoice" || q.Get("status") != "Open" {
			t.Errorf("missing filters on request %s", req.URL.String())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		pageNumber, _ := strconv.Atoi(q.Get("pageNumber"))
		if pageNumber == 0 {
			pageNumber = 1
		}
		pageSize, _ := strconv.Atoi(q.Get("pageSize"))
		if pageSize == 0 || pageSize > maxPageSize {
			pageSize = maxPageSize
		}

		start := (pageNumber - 1) * pageSize
		end := start + pageSize
		if end > totalCount {
			end = totalCount
		}

		invoices := []map[string]interface{}{}
		for i := start; i < end; i++ {
			invoices = append(invoices, map[string]interface{}{
				"referenceNumber": fmt.Sprintf("INV-%d", i+1),
				"metadata":        map[string]int{"totalCount": totalCount, "maxPageSize": maxPageSize},
			})
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(invoices)
	}))
	defer server.Close()

	srvURL, _ := url.Parse(server.URL)
	c := NewClient(nil)
	c.BaseURL = url.URL{Scheme: srvURL.Scheme, Host: srvURL.Host, Path: "/"}

	assertAll := func(name string, params GetCustomerInvoiceV1QueryParams) {
		atomic.StoreInt32(&requests, 0)

		req := c.NewGetCustomerInvoiceV1Request()
		req.SetQueryParams(params)
		resp, err := req.DoAll()
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}

		if len(resp.Invoices) != totalCount {
			t.Fatalf("%s: expected %d invoices, got %d", name, totalCount, len(resp.Invoices))
		}
		for i, inv := range resp.Invoices {
			if want := fmt.Sprintf("INV-%d", i+1); inv.ReferenceNumber != want {
				t.Errorf("%s: invoice %d: expected %s, got %s", name, i, want, inv.ReferenceNumber)
			}
		}

		// 1 discovery request + 2 concurrent page fetches = 3; no page is fetched twice.
		if got := atomic.LoadInt32(&requests); got != 3 {
			t.Errorf("%s: expected 3 requests, got %d", name, got)
		}
	}

	// Page size unset: fetches at the default page size (== maxPageSize).
	assertAll("default page size", GetCustomerInvoiceV1QueryParams{DocumentType: "Invoice", Status: "Open"})
	// A caller-set page size is ignored: pageSize 2 would need 4 requests if honored, but DoAll
	// still makes 3 because it always uses the default.
	assertAll("page size ignored", GetCustomerInvoiceV1QueryParams{DocumentType: "Invoice", Status: "Open", PageSize: 2})
}

// TestGetCustomerInvoiceV1DoAllConcurrencyBound verifies DoAll never runs more than
// Client.Concurrency page fetches at once, regardless of how many pages there are. The handler
// tracks concurrent in-flight requests and records the high-water mark; a small delay widens the
// window so overlapping requests are observed. With 10 pages and Concurrency 2, an unbounded
// fan-out would peak near 9; the bounded pool must keep the peak at Concurrency.
func TestGetCustomerInvoiceV1DoAllConcurrencyBound(t *testing.T) {
	const totalCount, maxPageSize, concurrency = 20, 2, 2

	var inFlight, maxInFlight int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Record the high-water mark of concurrent requests.
		cur := atomic.AddInt32(&inFlight, 1)
		for {
			old := atomic.LoadInt32(&maxInFlight)
			if cur <= old || atomic.CompareAndSwapInt32(&maxInFlight, old, cur) {
				break
			}
		}
		time.Sleep(10 * time.Millisecond) // widen the overlap window
		defer atomic.AddInt32(&inFlight, -1)

		q := req.URL.Query()
		pageNumber, _ := strconv.Atoi(q.Get("pageNumber"))
		if pageNumber == 0 {
			pageNumber = 1
		}
		pageSize, _ := strconv.Atoi(q.Get("pageSize"))
		if pageSize == 0 || pageSize > maxPageSize {
			pageSize = maxPageSize
		}

		start := (pageNumber - 1) * pageSize
		end := start + pageSize
		if end > totalCount {
			end = totalCount
		}

		invoices := []map[string]interface{}{}
		for i := start; i < end; i++ {
			invoices = append(invoices, map[string]interface{}{
				"referenceNumber": fmt.Sprintf("INV-%d", i+1),
				"metadata":        map[string]int{"totalCount": totalCount, "maxPageSize": maxPageSize},
			})
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(invoices)
	}))
	defer server.Close()

	srvURL, _ := url.Parse(server.URL)
	c := NewClient(nil)
	c.BaseURL = url.URL{Scheme: srvURL.Scheme, Host: srvURL.Host, Path: "/"}
	c.Concurrency = concurrency

	req := c.NewGetCustomerInvoiceV1Request()
	resp, err := req.DoAll()
	if err != nil {
		t.Fatal(err)
	}

	// All pages still fetched exactly once, in order.
	if len(resp.Invoices) != totalCount {
		t.Fatalf("expected %d invoices, got %d", totalCount, len(resp.Invoices))
	}
	for i, inv := range resp.Invoices {
		if want := fmt.Sprintf("INV-%d", i+1); inv.ReferenceNumber != want {
			t.Errorf("invoice %d: expected %s, got %s", i, want, inv.ReferenceNumber)
		}
	}

	// The discovery request completes before fan-out begins, so the peak should equal Concurrency;
	// allow +1 as slack against any scheduling overlap.
	if got := atomic.LoadInt32(&maxInFlight); got > concurrency+1 {
		t.Errorf("peak concurrent requests %d exceeded Concurrency %d (+1 slack)", got, concurrency)
	}
}

func TestGetCustomerInvoiceV1(t *testing.T) {
	invoiceNumber := os.Getenv("TEST_CUSTOMER_INVOICE_NUMBER")
	req := testClient.NewGetCustomerInvoiceV1Request()
	req.SetPathParams(GetCustomerInvoiceV1PathParams{invoiceNumber})
	resp, err := req.Do()
	debugDumpResponse(testClient, resp)
	if err != nil {
		t.Error(err)
	} else if resp.Invoices[0].ReferenceNumber != invoiceNumber {
		t.Errorf("Expected invoice number %s, got %s", invoiceNumber, resp.Invoices[0].ReferenceNumber)
	}
}

func TestPostCustomerInvoiceV2(t *testing.T) {
	req := testClient.NewPostCustomerInvoiceV2Request()
	req.SetBody(RequestInvoice{
		CurrencyID:     NewStringValue("SEK"),
		CustomerNumber: NewStringValue(os.Getenv("TEST_CUSTOMER_CD")),
		InvoiceLines: &[]RequestInvoiceLine{
			{
				Operation:           "Insert",
				ItemType:            NewStringValue("Service"),
				BranchNumber:        NewStringValue("1"),
				Description:         NewStringValue("Test"),
				Quantity:            NewFloatValue(1),
				UnitPriceInCurrency: NewFloatValue(100),
				AccountNumber:       NewStringValue("3015"),
				Subaccount: &[]RequestSegment{
					{
						SegmentID:    1,
						SegmentValue: "00",
					},
					{
						SegmentID:    2,
						SegmentValue: "000",
					},
					{
						SegmentID:    3,
						SegmentValue: "00",
					},
				},
			},
		},
	})
	resp, err := req.Do()
	debugDumpResponse(testClient, resp)
	if err != nil {
		t.Error(err)
	} else if resp.ResourceID() == "" {
		t.Errorf("Expected non-empty resource ID, got %s", resp.ResourceID())
	}
}

func TestDeleteCustomerInvoiceV1(t *testing.T) {
	req := testClient.NewDeleteCustomerInvoiceV1Request()
	req.SetPathParams(DeleteCustomerInvoiceV1PathParams{"1"})
	resp, err := req.Do()
	debugDumpResponse(testClient, resp)
	if err != nil && resp.StatusCode() != http.StatusNotFound {
		t.Error(err)
	}
}

func TestPostCustomerInvoiceAttachmentV1(t *testing.T) {
	req := testClient.NewPostCustomerInvoiceAttachmentV1Request()
	req.SetPathParams(PostCustomerInvoiceAttachmentV1PathParams{
		InvoiceNumber: os.Getenv("TEST_CUSTOMER_INVOICE_NUMBER"),
	})
	req.SetBody(FileUploadBody{
		Files: []File{
			{
				Name:    "Test.txt",
				Content: []byte("test"),
			},
		},
	})
	resp, err := req.Do()
	debugDumpResponse(testClient, resp)
	if err != nil {
		t.Error(err)
	} else if resp.ResourceID() == "" {
		t.Errorf("Expected non-empty resource ID, got %s", resp.ResourceID())
	}
}
