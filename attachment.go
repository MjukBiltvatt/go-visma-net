package vismanet

import "encoding/json"

type Attachment struct {
	ID       StringValue `json:"id,omitempty"`
	Name     StringValue `json:"name,omitempty"`
	Revision IntValue    `json:"revision,omitempty"`
}

// MarshalJSON wraps the Attachment in a Value struct and marshals it into a JSON byte slice
func (v *Attachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{*v})
}
