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
	if r.Http == nil {
		return ""
	}
	return r.Http.Status
}

// StatusCode returns the http status code
func (r *Response) StatusCode() int {
	if r.Http == nil {
		return 0
	}
	return r.Http.StatusCode
}

// Header returns the http header
func (r *Response) Header() http.Header {
	if r.Http == nil {
		return nil
	}
	return r.Http.Header
}

// GetHeader returns the value of the specified header
func (r *Response) GetHeader(key string) string {
	if r.Http == nil {
		return ""
	}
	return r.Http.Header.Get(key)
}

// LocationHeader returns the value of the 'Location' header
func (r *Response) LocationHeader() string {
	return r.GetHeader("Location")
}

// IPPRequestIDHeader returns the value of the 'Ipp-Request-Id' header
func (r *Response) IPPRequestIDHeader() string {
	return r.GetHeader("Ipp-Request-Id")
}

// RequestContextHeader returns the value of the 'Request-Context' header
func (r *Response) RequestContextHeader() string {
	return r.GetHeader("Request-Context")
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
