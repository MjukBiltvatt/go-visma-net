package vismanet

type Contact struct {
	ContactID int    `json:"contactId"`
	Name      string `json:"name"`
	Attention string `json:"attention"`
	Email     string `json:"email"`
	Web       string `json:"web"`
	Phone1    string `json:"phone1"`
	Phone2    string `json:"phone2"`
	Fax       string `json:"fax"`
}
