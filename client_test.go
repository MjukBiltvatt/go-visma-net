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
	closes   int
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
	b.closes++
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

// captureStdout redirects os.Stdout to a temporary file for the rest of the test, so debug dumps
// don't clutter the test output. The returned function reads what has been written so far.
func captureStdout(t *testing.T) func() string {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "stdout")
	if err != nil {
		t.Fatalf("failed to create stdout capture file: %v", err)
	}
	stdout := os.Stdout
	os.Stdout = f
	t.Cleanup(func() {
		os.Stdout = stdout
		f.Close()
	})
	return func() string {
		b, err := os.ReadFile(f.Name())
		if err != nil {
			t.Fatalf("failed to read captured stdout: %v", err)
		}
		return string(b)
	}
}

func TestDoClosesResponseBody(t *testing.T) {
	// The body of the response handed to Do must be closed exactly once, whatever the status code
	// is, whether or not reading it succeeds and whether or not debugging is enabled. net/http
	// requires it, and it is what disarms the http.Client.Timeout timer. The returned response must
	// behave like the drained body Do has always returned: it reads as empty and closes cleanly.
	cases := []struct {
		name    string
		status  int
		body    string
		readErr error
		debug   bool
		wantErr bool
	}{
		{name: "success", status: http.StatusOK, body: `{"referenceNumber":"INV-1"}`},
		{name: "error status", status: http.StatusBadRequest, body: `{"Message":"bad request"}`, wantErr: true},
		{name: "read error", status: http.StatusOK, body: `{}`, readErr: io.ErrUnexpectedEOF, wantErr: true},
		{name: "debug success", status: http.StatusOK, body: `{"referenceNumber":"INV-1"}`, debug: true},
		{name: "debug error status", status: http.StatusBadRequest, body: `{"Message":"bad request"}`, debug: true, wantErr: true},
		{name: "debug read error", status: http.StatusOK, body: `{}`, readErr: io.ErrUnexpectedEOF, debug: true, wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			body := &trackedBody{Reader: bytes.NewReader([]byte(tc.body)), readErr: tc.readErr}
			c := newStubClient(tc.status, body)
			c.Debug = tc.debug
			c.DebugBody = true
			stdout := captureStdout(t)

			req := c.NewGetCustomerInvoiceV1Request()
			req.SetPathParams(GetCustomerInvoiceV1PathParams{InvoiceNumber: "12345"})
			var invoice ResponseInvoice
			resp, err := c.Do((*Request)(&req), &invoice)
			if tc.wantErr && err == nil {
				t.Fatal("expected an error, got nil")
			} else if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if body.closes != 1 {
				t.Errorf("expected the original response body to be closed once, got %d closes", body.closes)
			}
			if resp == nil {
				t.Fatal("expected a response, got nil")
			}

			// The returned body must not be the (now closed) original, and must not hold on to the
			// body bytes.
			got, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("failed to read the returned response body: %v", err)
			}
			if len(got) != 0 {
				t.Errorf("expected the returned response body to be empty, got %q", string(got))
			}
			if err := resp.Body.Close(); err != nil {
				t.Errorf("failed to close the returned response body: %v", err)
			}

			// The debug dump must still include the response body, even though it has been closed.
			if tc.debug && tc.readErr == nil && !strings.Contains(stdout(), tc.body) {
				t.Errorf("expected the debug output to contain the response body %q", tc.body)
			}
		})
	}
}
