package vismanet

import (
	"bytes"
	"encoding/json"
	"io"
	"net/url"
	"reflect"
	"text/template"
)

// RequestBody represents a request body
type RequestBody interface {
	ContentType() string
	Reader() (io.Reader, error)
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

// URL returns the complete URL of the request
func (r *Request) URL() (string, error) {
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

// JSONRequestBody represents a request body that will be marshaled to JSON
type JSONRequestBody struct {
	// The value to be marshaled to JSON and sent as the request body
	Payload interface{}
}

// ContentType returns the content type of the request body
func (r JSONRequestBody) ContentType() string {
	return "application/json; charset=utf-8"
}

// Reader marshals the payload to JSON and returns a reader to be used as the request body
func (r JSONRequestBody) Reader() (io.Reader, error) {
	buf := new(bytes.Buffer)
	if r.Payload != nil {
		err := json.NewEncoder(buf).Encode(r.Payload)
		if err != nil {
			return nil, err
		}
	}
	return buf, nil
}
