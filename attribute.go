package vismanet

import "encoding/json"

type Attribute struct {
}

// MarshalJSON wraps the Attribute in a Value struct and marshals it into a JSON byte slice
func (v *Attribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{*v})
}
