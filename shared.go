package vismanet

type IDDescriptionEntity struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type NumberNameEntity struct {
	Number string `json:"number"`
	Name   string `json:"name"`
}

type IDNameEntity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Metadata struct {
	TotalCount  int `json:"totalCount"`
	MaxPageSize int `json:"maxPageSize"`
}
