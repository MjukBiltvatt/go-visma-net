package vismanet

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"golang.org/x/oauth2/clientcredentials"
)

var testClient *Client

func TestMain(m *testing.M) {
	oauthConf := clientcredentials.Config{
		ClientID:     os.Getenv("TEST_CLIENT_ID"),
		ClientSecret: os.Getenv("TEST_CLIENT_SECRET"),
		TokenURL:     os.Getenv("TEST_TOKEN_URL"),
		Scopes:       strings.Split(os.Getenv("TEST_SCOPES"), " "),
		EndpointParams: map[string][]string{
			"tenant_id": {os.Getenv("TEST_TENANT_ID")},
		},
	}
	httpClient := oauthConf.Client(context.Background())
	testClient = NewClient(httpClient)
	testClient.Debug = os.Getenv("DEBUG") == "true"
	testClient.DebugBody = os.Getenv("DEBUG_BODY") != "false"
	m.Run()
}

// trackedBody is a response body that records whether it has been closed.
type trackedBody struct {
	io.Reader
	closed   bool
	readErr  error
	readSome bool
}

func (b *trackedBody) Read(p []byte) (int, error) {
	if b.readErr != nil {
		// Hand out some bytes before failing, so partial reads are exercised too.
		if !b.readSome {
			b.readSome = true
			return copy(p, "partial"), nil
		}
		return 0, b.readErr
	}
	return b.Reader.Read(p)
}

func (b *trackedBody) Close() error {
	b.closed = true
	return nil
}

// roundTripperFunc adapts a function to the http.RoundTripper interface.
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

// newStubClient returns a client whose transport always responds with the given status and body.
func newStubClient(status int, body *trackedBody) *Client {
	return NewClient(&http.Client{
		Transport: roundTripperFunc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				Status:     http.StatusText(status),
				StatusCode: status,
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				Body:       body,
				Request:    r,
			}, nil
		}),
	})
}

func TestDoClosesResponseBody(t *testing.T) {
	// The body of the response handed to Do must always be closed, whatever the status code is and
	// whether or not reading it succeeds. net/http requires it, and it is what disarms the
	// http.Client.Timeout timer. The returned response must still expose a readable body, since Do
	// returns the response to its callers.
	cases := []struct {
		name    string
		status  int
		body    string
		readErr error
		wantErr bool
	}{
		{name: "success", status: http.StatusOK, body: `{"referenceNumber":"INV-1"}`},
		{name: "error status", status: http.StatusBadRequest, body: `{"Message":"bad request"}`, wantErr: true},
		{name: "read error", status: http.StatusOK, body: `{}`, readErr: io.ErrUnexpectedEOF, wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			body := &trackedBody{Reader: bytes.NewReader([]byte(tc.body)), readErr: tc.readErr}
			c := newStubClient(tc.status, body)

			req := c.NewGetCustomerInvoiceV1Request()
			req.SetPathParams(GetCustomerInvoiceV1PathParams{InvoiceNumber: "12345"})
			var invoice ResponseInvoice
			resp, err := c.Do((*Request)(&req), &invoice)
			if tc.wantErr && err == nil {
				t.Fatal("expected an error, got nil")
			} else if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !body.closed {
				t.Error("expected the original response body to be closed")
			}
			if resp == nil {
				t.Fatal("expected a response, got nil")
			}

			// The returned body must be readable, and must not be the (now closed) original.
			got, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("failed to read the returned response body: %v", err)
			}
			if tc.readErr == nil && string(got) != tc.body {
				t.Errorf("expected returned body %q, got %q", tc.body, string(got))
			}
			if err := resp.Body.Close(); err != nil {
				t.Errorf("failed to close the returned response body: %v", err)
			}
		})
	}
}

func TestDoClosesResponseBodyWhenDebugging(t *testing.T) {
	// Debug dumping replaces resp.Body with a copy of its own, so the buffering and closing in Do
	// has to keep working when it is enabled.
	body := &trackedBody{Reader: bytes.NewReader([]byte(`{"referenceNumber":"INV-1"}`))}
	c := newStubClient(http.StatusOK, body)
	c.Debug = true
	c.DebugBody = true

	req := c.NewGetCustomerInvoiceV1Request()
	req.SetPathParams(GetCustomerInvoiceV1PathParams{InvoiceNumber: "12345"})
	var invoice ResponseInvoice
	resp, err := c.Do((*Request)(&req), &invoice)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !body.closed {
		t.Error("expected the original response body to be closed")
	}
	if got, err := io.ReadAll(resp.Body); err != nil {
		t.Fatalf("failed to read the returned response body: %v", err)
	} else if len(got) == 0 {
		t.Error("expected the returned response body to be readable")
	}
}
