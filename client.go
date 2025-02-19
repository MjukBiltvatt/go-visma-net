package vismanet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// NewClient returns a new Visma Net client either using the specified http.Client, or if nil the default http.Client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		Http: httpClient,
		BaseURL: url.URL{
			Scheme: "https",
			Host:   "integration.visma.net",
			Path:   "/API/",
		},
		Debug:     false,
		UserAgent: UserAgent,
		Charset:   "utf-8",
	}
}

// Client is a Visma Net client for making Visma Net API requests.
type Client struct {
	Http      *http.Client
	BaseURL   url.URL
	Debug     bool
	UserAgent string
	Charset   string
}

// Do the API request and decode the response body into the provided interface
func (c *Client) Do(req *Request, body ...interface{}) (*http.Response, error) {
	//Get complete URL
	url, err := req.URL()
	if err != nil {
		return nil, err
	}

	//Get body reader
	var reqBody io.Reader
	if req.Body != nil {
		reqBody, err = req.Body.Reader()
		if err != nil {
			return nil, err
		}
	}

	//Create http request
	r, err := http.NewRequest(req.Method, url, reqBody)
	if err != nil {
		return nil, err
	}

	//Add headers
	r.Header.Add("Accept", "application/json")
	r.Header.Add("User-Agent", c.UserAgent)
	if req.Body != nil {
		r.Header.Add("Content-Type", req.Body.ContentType())
	}

	//Execute http request
	resp, err := c.Http.Do(r)
	if err != nil {
		return resp, err
	}

	//Dump response if debugging is enabled
	if c.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, fmt.Errorf("failed to dump response: %v", err)
		}
		fmt.Println(string(dump))
	}

	//Check response status code
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		//Decode error response
		var errResp ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil && err != io.EOF {
			return resp, fmt.Errorf("failed to decode error response: %v", err)
		}
		return resp, errResp
	}

	// Read the response body into a buffer
	var buffer bytes.Buffer
	if _, err := io.Copy(&buffer, resp.Body); err != nil {
		return resp, fmt.Errorf("failed to read response body: %v", err)
	}

	// Attempt to decode the response body into any of the provided interfaces
	var decodeErr error
	for _, b := range body {
		// Create a new buffer for each decode attempt
		var buf bytes.Buffer
		if _, err := buf.Write(buffer.Bytes()); err != nil {
			return resp, fmt.Errorf("failed to write response body to buffer: %v", err)
		}

		// Attempt to decode the response body into the provided interface
		if err := json.NewDecoder(&buf).Decode(b); err != nil && err != io.EOF {
			decodeErr = fmt.Errorf("failed to decode response body: %v", err)
		} else {
			return resp, nil
		}
	}

	// Failed to decode the response body into any of the provided interfaces, return error
	return resp, decodeErr
}

func (c *Client) NewGetCustomerV1Request() GetCustomerV1Request {
	return newGetCustomerV1Request(c)
}

func (c *Client) NewPostCustomerV1Request() PostCustomerV1Request {
	return newPostCustomerV1Request(c)
}
