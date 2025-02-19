package vismanet

type Contact struct {
	ContactID int         `json:"contactId,omitempty"`
	Name      StringValue `json:"name,omitempty"`
	Attention StringValue `json:"attention,omitempty"`
	Email     StringValue `json:"email,omitempty"`
	Web       StringValue `json:"web,omitempty"`
	Phone1    StringValue `json:"phone1,omitempty"`
	Phone2    StringValue `json:"phone2,omitempty"`
	Fax       StringValue `json:"fax,omitempty"`
}
