package vismanet

type Address struct {
	AddressID    int     `json:"addressId"`
	AddressLine1 string  `json:"addressLine1"`
	AddressLine2 string  `json:"addressLine2"`
	AddressLine3 string  `json:"addressLine3"`
	PostalCode   string  `json:"postalCode"`
	City         string  `json:"city"`
	Country      Country `json:"country"`
}
