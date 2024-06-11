package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Shipping struct {
	ID           string             `json:"_id,omitempty"  bson:"_id,omitempty"`
	OrderId      primitive.ObjectID `json:"order_id" bson:"order_id"`
	ShippingDate primitive.DateTime `json:"shipping_date" bson:"shipping_date"`
	DeliveryDate primitive.DateTime `json: "delivery_date" bson:"delivery_date"`
	Status       string             `json:"status" bson:"status"`
}
