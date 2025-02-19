package vismanet

import "encoding/json"

// Contact is a contact as represented in responses from Visma.net
type Contact struct {
	ContactID int    `json:"contactId"`
	Name      string `json:"name"`
	Attention string `json:"attention"`
	Email     string `json:"email"`
	Web       string `json:"web"`
	Phone1    string `json:"phone1"`
	Phone2    string `json:"phone2"`
	Fax       string `json:"fax"`
}

// RequestBodyNestedContact is a contact nested in another entity in the body of requests to Visma.net
type RequestBodyNestedContact struct {
	ContactID int         `json:"contactId,omitempty"`
	Name      StringValue `json:"name,omitempty"`
	Attention StringValue `json:"attention,omitempty"`
	Email     StringValue `json:"email,omitempty"`
	Web       StringValue `json:"web,omitempty"`
	Phone1    StringValue `json:"phone1,omitempty"`
	Phone2    StringValue `json:"phone2,omitempty"`
	Fax       StringValue `json:"fax,omitempty"`
}

// MarshalJSON wraps the Contact in a Value struct and marshals it into a JSON byte slice
func (v *RequestBodyNestedContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{*v})
}
