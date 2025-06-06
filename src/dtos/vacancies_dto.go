package dtos

type Job struct {
	Id          string `json:"id"`
	Titulo      string `json:"titulo"`
	Descricao   string `json:"descricao"`
	Localizacao string `json:"localizacao"`
	PostedAt    string `json:"posted:"`
}
