package vismanet

type Account struct {
	Type        StringValue `json:"type"`
	Number      StringValue `json:"number"`
	Description StringValue `json:"description"`
}
