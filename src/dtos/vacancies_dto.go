package dtos

type Vacancies struct {
	Id           string   `json:"id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Area         string   `json:"area"`
	Category     string   `json:"type"`
	Location     string   `json:"location"`
	Requirements []string `json:"requirements"`
	PostedAt     string   `json:"posted:"`
}
