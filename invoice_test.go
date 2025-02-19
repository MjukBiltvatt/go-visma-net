package vismanet

import (
	"os"
	"testing"
)

func TestGetInvoiceV1(t *testing.T) {
	invoiceNumber := os.Getenv("TEST_INVOICE_NUMBER")
	req := testClient.NewGetInvoiceV1Request()
	req.SetPathParams(GetInvoiceV1PathParams{invoiceNumber})
	resp, err := req.Do()
	debugDumpResponse(testClient, resp)
	if err != nil {
		t.Error(err)
	} else if resp.Invoices[0].ReferenceNumber != invoiceNumber {
		t.Errorf("Expected invoice number %s, got %s", invoiceNumber, resp.Invoices[0].ReferenceNumber)
	}
}

func TestPostInvoiceV2(t *testing.T) {
	req := testClient.NewPostInvoiceV2Request()
	req.SetBody(RequestInvoice{
		CurrencyID:     "SEK",
		CustomerNumber: StringValue(os.Getenv("TEST_CUSTOMER_CD")),
		InvoiceLines: []RequestInvoiceLine{
			{
				Operation:           "Insert",
				ItemType:            "Service",
				BranchNumber:        "1",
				Description:         "Test",
				Quantity:            1,
				UnitPriceInCurrency: 100,
				AccountNumber:       "3015",
				Subaccount: []RequestSegment{
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
