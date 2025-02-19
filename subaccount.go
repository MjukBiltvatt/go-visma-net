package vismanet

import "encoding/json"

type Subaccount struct {
	SubaccountID         IntValue    `json:"subaccountId"`
	SubaccountNumber     StringValue `json:"subaccountNumber"`
	Description          StringValue `json:"description,omitempty"`
	LastModifiedDateTime Time        `json:"lastModifiedDateTime,omitempty"`
	Active               BoolValue   `json:"active"`
	Segments             []Segment   `json:"segments"`
}

// MarshalJSON wraps the Subaccount in a Value struct and marshals it into a JSON byte slice
func (v *Subaccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{*v})
}
