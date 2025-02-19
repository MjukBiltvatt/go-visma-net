package vismanet

import "encoding/json"

type Contact struct {
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
func (v *Contact) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{*v})
}
