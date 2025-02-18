package vismanet

type Subaccount struct {
	SubaccountID         int       `json:"subaccountId"`
	SubaccountNumber     string    `json:"subaccountNumber"`
	Description          string    `json:"description"`
	LastModifiedDateTime Time      `json:"lastModifiedDateTime"`
	Active               bool      `json:"active"`
	Segments             []Segment `json:"segments"`
}
