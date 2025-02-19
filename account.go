package vismanet

import "encoding/json"

type Account struct {
	Type        StringValue `json:"type"`
	Number      StringValue `json:"number"`
	Description StringValue `json:"description"`
}

// MarshalJSON wraps the Account in a Value struct and marshals it into a JSON byte slice
func (v *Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{*v})
}
