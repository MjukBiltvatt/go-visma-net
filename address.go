package vismanet

import "encoding/json"

type Address struct {
	AddressID    int         `json:"addressId,omitempty"`
	AddressLine1 StringValue `json:"addressLine1,omitempty"`
	AddressLine2 StringValue `json:"addressLine2,omitempty"`
	AddressLine3 StringValue `json:"addressLine3,omitempty"`
	PostalCode   StringValue `json:"postalCode,omitempty"`
	City         StringValue `json:"city,omitempty"`
	CountryID    StringValue `json:"countryId,omitempty"`
	County       StringValue `json:"county,omitempty"`
}

// MarshalJSON wraps the Address in a Value struct and marshals it into a JSON byte slice
func (v *Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{*v})
}
