package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Job struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"ID"`
	Titulo    string             `bson:"titulo" json:"Titulo"`
	Descricao string             `bson:"descricao" json:"Descricao"`
	Local     string             `bson:"local" json:"Local"`
	PostedAt  string             `bson:"postedAt" json:"PostedAt"`
}
