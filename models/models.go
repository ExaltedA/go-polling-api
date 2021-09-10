package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Question string             `json:"question,omitempty"`
	Variant_1 string             `json:"variant_1,omitempty"`
	Variant_2 string             `json:"variant_2,omitempty"`
	Variant_3 string             `json:"variant_3,omitempty"`
	Variant_4 string             `json:"variant_4,omitempty"`
	Right_answer int             `json:"right_answer,omitempty"`
}
