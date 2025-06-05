package dtos

type Curriculo struct {
	Nome        string      `json:"nome"`
	Sobrenome   string      `json:"sobrenome"`
	Idade       int         `json:"idade"`
	Email       string      `json:"email"`
	Formacao    []Education `json:"formacao"`
	Experiencia []string    `json:"experiencia"`
}
