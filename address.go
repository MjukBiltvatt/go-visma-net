package vismanet

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
