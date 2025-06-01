package dtos

import "time"

// Job representa a estrutura de uma vaga de emprego.
type Job struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	PostedAt    time.Time `json:"posted:"`
}
