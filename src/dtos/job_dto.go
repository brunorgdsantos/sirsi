package dtos

type Job struct {
	Titulo      string `json:"titulo"`
	Descricao   string `json:"descricao"`
	Localizacao string `json:"localizacao"`
	PostedAt    string `json:"posted:"`
}
