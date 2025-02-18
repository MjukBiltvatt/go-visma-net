package vismanet

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
func (c *Client) Do(req *Request, body interface{}) (*http.Response, error) {
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

	//Check response status code
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		//Decode error response
		var errResp ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&err); err != nil {
			return resp, fmt.Errorf("failed to decode error response: %v", err)
		}
		return resp, errResp
	}

	//Decode response body
	if err := json.NewDecoder(resp.Body).Decode(body); err != nil {
		return resp, fmt.Errorf("failed to decode response body: %v", err)
	}

	return resp, nil
}
