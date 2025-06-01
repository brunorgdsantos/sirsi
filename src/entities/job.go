package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Job struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"ID"`
	Titulo    string             `bson:"titulo" json:"Titulo"`
	Descricao string             `bson:"descricao" json:"Descricao"`
	Local     string             `bson:"local" json:"Local"`
	PostedAt  time.Time          `bson:"postedAt" json:"PostedAt"`
}
