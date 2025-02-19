package vismanet

type Subaccount struct {
	SubaccountID         IntValue    `json:"subaccountId"`
	SubaccountNumber     StringValue `json:"subaccountNumber"`
	Description          StringValue `json:"description,omitempty"`
	LastModifiedDateTime TimeValue   `json:"lastModifiedDateTime,omitempty"`
	Active               BoolValue   `json:"active"`
	Segments             []Segment   `json:"segments"`
}
