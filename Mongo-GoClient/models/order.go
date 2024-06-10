package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         string             `json:"_id,omitempty"  bson:"_id,omitempty"`
	CustomerId primitive.ObjectID `json:"customer_id" bson:"customer_id"`
	OrderDate  primitive.DateTime `json:"order_date" bson:"order_date"`
	Status     string             `json:"status" bson:"status"`
	Items      []interface{}      `json:"items" bson:"items"`
	ShippingID primitive.ObjectID `json:"shipping_id" bson:"shipping_id"`
}
