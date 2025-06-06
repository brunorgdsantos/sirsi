package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vacancies struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"ID"`
	Title        string             `bson:"title" json:"Title"`
	Description  string             `bson:"description" json:"Description"`
	Area         string             `bson:"area" json:"Area"`
	Category     string             `bson:"type" json:"Type"`
	Location     string             `bson:"local" json:"Location"`
	Requirements []string           `bson:"requirements" json:"Requirements"`
	PostedAt     string             `bson:"postedAt" json:"PostedAt"`
}
