package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          string             `json:"_id,omitempty"  bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Category    string             `json:"category" bson:"category"`
	Price       int                `json:"price" bson:"price"`
	InventoryId primitive.ObjectID `json:"inventory_id" bson:"inventory_id"`
}
