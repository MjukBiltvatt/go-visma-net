package vismanet

import (
	"net/http"
	"path"
)

// Response represents a Visma Net API response
type Response struct {
	Http *http.Response
}

// StatusCode returns the http status
func (r *Response) Status() string {
	return r.Http.Status
}

// StatusCode returns the http status code
func (r *Response) StatusCode() int {
	return r.Http.StatusCode
}

// Header returns the http header
func (r *Response) Header() http.Header {
	return r.Http.Header
}

// LocationHeader returns the value of the 'Location' header
func (r *Response) LocationHeader() string {
	return r.Http.Header.Get("Location")
}

// ResourceID returns the resource id from the 'Location' header
func (r *Response) ResourceID() string {
	return path.Base(r.LocationHeader())
}

// ErrorResponse represents a Visma Net API error response
type ErrorResponse struct {
	// HTTP response that caused the error
	Response *http.Response
	Message  string `json:"Message"`
}

// Error returns the error message
func (r ErrorResponse) Error() string {
	return r.Message
}
