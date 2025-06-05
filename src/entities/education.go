package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Education struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"ID"`
	institution string             `bson:"institution" json:"institution"`
	course      string             `bson:"course" json:"course"`
	level       string             `bson:"level" json:"level"`
}
