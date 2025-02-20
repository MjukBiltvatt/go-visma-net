package vismanet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"reflect"
	"text/template"
)

// RequestBody represents a request body
type RequestBody interface {
	build() (io.Reader, string, error)
}

// Request represents an API request
type Request struct {
	Client      *Client
	Method      string
	Path        string
	Body        RequestBody
	queryParams url.Values
	pathParams  interface{}
}

// pathParamsMap returns a map of path parameters name/value pairs from the pathParams struct
func (r *Request) pathParamsMap() map[string]string {
	if r.pathParams == nil {
		return map[string]string{}
	}

	//Create map of path parameters name/value pairs
	params := make(map[string]string)

	//Get the type and value of the path parameters struct
	t := reflect.TypeOf(r.pathParams)
	v := reflect.ValueOf(r.pathParams)

	//Iterate over the fields of the path parameters struct and add any string fields with a schema tag to the map
	for i := 0; i < t.NumField(); i++ {
		name, ok := t.Field(i).Tag.Lookup("schema")
		if !ok || t.Field(i).Type.Kind() != reflect.String {
			continue
		}
		params[name] = v.Field(i).String()
	}

	return params
}

// url returns the complete URL of the request
func (r *Request) url() (string, error) {
	//Parse path template
	tmpl, err := template.New("path").Option("missingkey=error").Parse(r.Path)
	if err != nil {
		return "", err
	}

	//Execute path template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, r.pathParamsMap()); err != nil {
		return "", err
	}

	//Return the complete URL
	return r.Client.BaseURL.String() + buf.String(), nil
}

// build the http request
func (r *Request) build() (*http.Request, error) {
	// Get complete URL
	url, err := r.url()
	if err != nil {
		return nil, err
	}

	// Get request body reader
	var reqBody io.Reader
	var contentType string
	if r.Body != nil {
		reqBody, contentType, err = r.Body.build()
		if err != nil {
			return nil, err
		}
	}

	// Create http request
	req, err := http.NewRequest(r.Method, url, reqBody)
	if err != nil {
		return nil, err
	}

	// Add headers
	req.Header.Add("User-Agent", r.Client.UserAgent)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Accept", "application/json")

	return req, nil
}

// JSONRequestBody represents a request body that will be marshaled to JSON
type JSONRequestBody struct {
	// The value to be marshaled to JSON and sent as the request body
	Payload interface{}
}

// build marshals the payload to JSON and returns a reader to be used as the request body as well as the content type
func (r JSONRequestBody) build() (io.Reader, string, error) {
	contentType := "application/json; charset=utf-8"
	buf := new(bytes.Buffer)
	if r.Payload != nil {
		err := json.NewEncoder(buf).Encode(r.Payload)
		if err != nil {
			return nil, contentType, err
		}
	}
	return buf, contentType, nil
}

type File struct {
	Key         string // The key to use for the file in the multipart form data
	Name        string // The name of the file
	Content     []byte // The content of the file
	ContentType string // The content type of the file, default "application/octet-stream" if left empty
}

// FileUploadBody represents a multipart form data request body builder
type FileUploadBody struct {
	Files []File
}

// AddFile adds a file to the request body
func (r *FileUploadBody) AddFile(file File) {
	r.Files = append(r.Files, file)
}

// build marshals the payload to JSON and returns a reader to be used as the request body
func (r FileUploadBody) build() (io.Reader, string, error) {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	defer w.Close()

	for _, file := range r.Files {
		// Create the form part header
		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, file.Key, file.Name))
		if file.ContentType != "" {
			h.Set("Content-Type", file.ContentType)
		} else {
			h.Set("Content-Type", "application/octet-stream")
		}

		// Create the form part
		part, err := w.CreatePart(h)
		if err != nil {
			return nil, w.FormDataContentType(), fmt.Errorf("failed to create form part: %v", err.Error())
		}

		// Write the file data to the form part
		if _, err = part.Write(file.Content); err != nil {
			return nil, w.FormDataContentType(), fmt.Errorf("failed to write to form part: %v", err.Error())
		}
	}

	return buf, w.FormDataContentType(), nil
}
