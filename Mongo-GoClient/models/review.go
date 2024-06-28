package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	ID         string             `json:"_id,omitempty"  bson:"_id,omitempty"`
	ProductId  primitive.ObjectID `json:"product_id" bson:"product_id"`
	CustomerId primitive.ObjectID `json:"customer_id" bson:"customer_id"`
	Rating     float64            `json:"rating" bson:"rating"`
	ReviewDate primitive.DateTime `json:"review_date" bson:"review_date"`
}
