package vismanet

import "encoding/json"

type Location struct {
	ID        string `json:"id"`
	CountryID string `json:"countryId"`
	Name      string `json:"name"`
}

// RequestLocation is a location as represented in a request to the Visma.net API
type RequestLocation struct {
	BAccountID          *StringValue                 `json:"baccountId,omitempty"`
	LocationID          *StringValue                 `json:"locationId,omitempty"`
	LocationName        *StringValue                 `json:"locationName,omitempty"`
	Active              *BoolValue                   `json:"active,omitempty"`
	AddressIsSameAsMain *BoolValue                   `json:"addressIsSameAsMain,omitempty"`
	Address             *RequestNestedAddress        `json:"address,omitempty"`
	ContactIsSameAsMain *BoolValue                   `json:"contactIsSameAsMain,omitempty"`
	Contact             *RequestNestedContact        `json:"contact,omitempty"`
	VatRegistrationID   *StringValue                 `json:"vatRegistrationId,omitempty"`
	VatZone             *StringValue                 `json:"vatZone,omitempty"`
	EdiCode             *StringValue                 `json:"ediCode,omitempty"`
	GLN                 *StringValue                 `json:"gln,omitempty"`
	CorporateID         *StringValue                 `json:"corporateId,omitempty"`
	PeppolScheme        *RequestLocationPeppolScheme `json:"peppolScheme,omitempty"`
}

// RequestLocationPeppolScheme is a location peppol scheme as represented in a request to the Visma.net API
type RequestLocationPeppolScheme struct {
	Endpoint            *StringValue `json:"endpoint,omitempty"`
	PartyIdentification *StringValue `json:"partyIdentification,omitempty"`
	PartyLegal          *StringValue `json:"partyLegal,omitempty"`
}

func (v *RequestLocationPeppolScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{*v})
}

// =========================================================
// ========================== PUT ==========================
// =========================================================

// newPutLocationV1Request creates a new PutLocationV1Request
func newPutLocationV1Request(c *Client) PutLocationV1Request {
	return PutLocationV1Request{
		Client: c,
		Method: "PUT",
		Path:   "controller/api/v1/location/{{.b_account_id}}/{{.location_id}}",
	}
}

// PutLocationV1Request represents a request to create a new customer
type PutLocationV1Request Request

// SetPathParams sets the path parameters of the request
func (r *PutLocationV1Request) SetPathParams(params PutLocationV1PathParams) {
	r.pathParams = params
}

// SetBody sets the body of the request
func (r *PutLocationV1Request) SetBody(body RequestLocation) {
	r.Body = JSONRequestBody{body}
}

// Do performs the request and returns the response
func (r *PutLocationV1Request) Do() (PutLocationV1Response, error) {
	resp, err := r.Client.Do((*Request)(r), nil)
	return PutLocationV1Response{Response{resp}}, err
}

// PutLocationV1PathParams represents the path parameters of the PutLocationV1Request
type PutLocationV1PathParams struct {
	BAccountID string `schema:"b_account_id"`
	LocationID string `schema:"location_id"`
}

// PutLocationV1Response represents the response of the PutLocationV1Request
type PutLocationV1Response struct {
	Response
}
