package vismanet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

// NewClient returns a new Visma Net client either using the specified http.Client, or if nil the default http.Client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	return &Client{
		Http: httpClient,
		BaseURL: url.URL{
			Scheme: "https",
			Host:   "integration.visma.net",
			Path:   "/API/",
		},
		Debug:     false,
		DebugBody: true,
		UserAgent: UserAgent,
		Charset:   "utf-8",
	}
}

// Client is a Visma Net client for making Visma Net API requests.
type Client struct {
	Http      *http.Client // Http is the http client used to make requests
	BaseURL   url.URL      // BaseURL is the base URL of the Visma Net API
	Debug     bool         // Debug enables debugging output of requests and responses
	DebugBody bool         // DebugBody enables debugging output of request and response bodies
	UserAgent string       // UserAgent is the string used in the User-Agent header in requests
	Charset   string       // Charset is the character set used in the Content-Type header in requests
}

// Do the API request and decode the response body into the provided interface
func (c *Client) Do(req *Request, body ...interface{}) (*http.Response, error) {
	// Build http request
	r, err := req.build()
	if err != nil {
		return nil, fmt.Errorf("failed to build http request: %w", err)
	}

	// Dump request if debugging is enabled
	if c.Debug {
		dump, err := httputil.DumpRequest(r, c.DebugBody)
		if err != nil {
			return nil, fmt.Errorf("failed to dump request: %w", err)
		}
		fmt.Println(string(dump))
	}

	// Execute http request
	resp, err := c.Http.Do(r)
	if err != nil {
		return resp, fmt.Errorf("http request failed: %w", err)
	}

	// Dump response if debugging is enabled
	if c.Debug {
		dump, err := httputil.DumpResponse(resp, c.DebugBody)
		if err != nil {
			return resp, fmt.Errorf("failed to dump response: %w", err)
		}
		fmt.Println(string(dump))
	}

	// Read the response body into a buffer
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return resp, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check response status code
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		// Decode error response
		var errResp ErrorResponse
		if err := json.Unmarshal(buf.Bytes(), &errResp); err != nil && err != io.EOF {
			return resp, fmt.Errorf("failed to decode error response: %w", err)
		} else if errResp.Message != "" {
			// Return error response as error
			return resp, errResp
		}

		// Decode exception response
		var excResp ExceptionResponse
		if err := json.Unmarshal(buf.Bytes(), &excResp); err != nil && err != io.EOF {
			return resp, fmt.Errorf("failed to decode exception response: %w", err)
		} else if excResp.Message != "" {
			// Return exception response as error
			return resp, excResp
		}

		// Return error with http status if no error or exception response could be decoded
		return resp, fmt.Errorf("http status %s indicates error, but the body could not be parsed: '%s'", resp.Status, buf.String())
	}

	// Attempt to decode the response body into any of the provided interfaces
	var decodeErr error
	for _, b := range body {
		// Skip nil interfaces
		if b == nil {
			continue
		}

		// Unmarshal response body into provided interface
		if err := json.Unmarshal(buf.Bytes(), b); err != nil && err != io.EOF {
			decodeErr = fmt.Errorf("failed to decode response body: %w", err)
		} else {
			return resp, nil
		}
	}

	// Failed to decode the response body into any of the provided interfaces, return error
	return resp, decodeErr
}

// NewGetCustomerV1Request creates a new GetCustomerV1Request
func (c *Client) NewGetCustomerV1Request() GetCustomerV1Request {
	return newGetCustomerV1Request(c)
}

// NewPostCustomerV1Request creates a new PostCustomerV1Request
func (c *Client) NewPostCustomerV1Request() PostCustomerV1Request {
	return newPostCustomerV1Request(c)
}

// NewPutCustomerV1Request creates a new PutCustomerV1Request
func (c *Client) NewPutCustomerV1Request() PutCustomerV1Request {
	return newPutCustomerV1Request(c)
}

// NewGetCustomerInvoiceV1Request creates a new GetCustomerInvoiceV1Request
func (c *Client) NewGetCustomerInvoiceV1Request() GetCustomerInvoiceV1Request {
	return newGetCustomerInvoiceV1Request(c)
}

// NewPostCustomerInvoiceV2Request creates a new PostCustomerInvoiceV2Request
func (c *Client) NewPostCustomerInvoiceV2Request() PostCustomerInvoiceV2Request {
	return newPostCustomerInvoiceV2Request(c)
}

// NewPostCustomerInvoiceAttachmentV1Request creates a new PostCustomerInvoiceAttachmentV1Request
func (c *Client) NewPostCustomerInvoiceAttachmentV1Request() PostCustomerInvoiceAttachmentV1Request {
	return newPostCustomerInvoiceAttachmentV1Request(c)
}

// NewDeleteCustomerInvoiceV1Request creates a new DeleteCustomerInvoiceV1Request
func (c *Client) NewDeleteCustomerInvoiceV1Request() DeleteCustomerInvoiceV1Request {
	return newDeleteCustomerInvoiceV1Request(c)
}
