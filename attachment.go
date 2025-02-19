package vismanet

// Attachment is an attachment as represented in responses from Visma.net
type Attachment struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Revision int    `json:"revision"`
}
