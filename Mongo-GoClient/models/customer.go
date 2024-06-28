package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	City string `json:"city" bson:"city"`
	Zip  string `json:"zip" bson:"zip"`
}
type Customer struct {
	ID       string             `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	JoinDate primitive.DateTime `json:"join_date" bson:"join_date"`
	Address  Address            `json:"address" bson:"address"`
}
