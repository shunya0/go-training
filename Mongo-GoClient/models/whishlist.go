package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type WhishlistItems struct {
	ProductId primitive.ObjectID `json:"product_id" bson:"product_id"`
}
type Whishlist struct {
	ID         string             `json:"_id,omitempty"  bson:"_id,omitempty"`
	CustomerId primitive.ObjectID `json:"customer_id" bson:"customer_id"`
	Item       []WhishlistItems   `json:"items" bson:"items"`
	AddedDate  primitive.DateTime `json:"added_date" bson:"added_date"`
}
