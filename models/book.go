package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Author   string             `json:"author,omitempty" bson:"author,omitempty"`
	Language string             `json:"language,omitempty" bson:"language,omitempty"`
	Pages    int32              `json:"pages,omitempty" bson:"pages,omitempty"`
	Price    float64            `json:"price,omitempty" bson:"price,omitempty"`
	Domain   string             `json:"domain,omitempty" bson:"domain,omitempty"`
}
