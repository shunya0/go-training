package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Inventory struct {
	ID                string             `json:"_id,omitempty"  bson:"_id,omitempty"`
	ProductId         primitive.ObjectID `json:"product_id" bson:"product_id"`
	QuantityAvailable int                `json:"quantity_available" bson:"quantity_available"`
}
