package vismanet

import "encoding/json"

// Address is an address as represented in responses from Visma.net
type Address struct {
	AddressID       int          `json:"addressId"`
	AddressLine1    string       `json:"addressLine1"`
	AddressLine2    string       `json:"addressLine2"`
	AddressLine3    string       `json:"addressLine3"`
	PostalCode      string       `json:"postalCode"`
	City            string       `json:"city"`
	OverrideAddress bool         `json:"overrideAddress"`
	Country         Country      `json:"country"`
	County          IDNameEntity `json:"county"`
}

// Country is a country as represented in responses from Visma.net
type Country struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	ErrorInfo string   `json:"errorInfo"`
	Metadata  Metadata `json:"metadata"`
}

// RequestBodyNestedAddress is an address nested in another entity in the body of requests to Visma.net
type RequestBodyNestedAddress struct {
	AddressLine1    StringValue `json:"addressLine1,omitempty"`
	AddressLine2    StringValue `json:"addressLine2,omitempty"`
	AddressLine3    StringValue `json:"addressLine3,omitempty"`
	PostalCode      StringValue `json:"postalCode,omitempty"`
	City            StringValue `json:"city,omitempty"`
	CountryID       StringValue `json:"countryId,omitempty"`
	County          StringValue `json:"county,omitempty"`
	OverrideAddress BoolValue   `json:"overrideAddress,omitempty"`
}

// MarshalJSON wraps the Address in a Value struct and marshals it into a JSON byte slice
func (v *RequestBodyNestedAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{*v})
}
