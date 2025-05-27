package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"ID"`
	Nome        string             `bson:"nome" json:"Nome"`
	Sobrenome   string             `bson:"sobrenome" json:"Sobrenome"`
	Idade       int                `bson:"idade" json:"idade"`
	Email       string             `bson:"email" json:"Email"`
	Experiencia int                `bson:"experiencia" json:"Experiencia"`
	Formacao    string             `bson:"formacao" json:"Formacao"`
	Password    string             `bson:"password" json:"-"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}
