package dtos

type Curriculo struct {
	Nome        string `json:"nome"`
	Sobrenome   string `json:"sobrenome"`
	Formacao    string `json:"formacao"`
	Experiencia string `json:"experiencia"`
}
