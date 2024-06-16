package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cancel struct {
	OrderID        primitive.ObjectID `json:"order_id" bson:"order_id"`
	CustomerID     primitive.ObjectID `json:"customer_id" bson:"customer_id"`
	ShippingID     primitive.ObjectID `json:"shipping_id" bson:"shipping_id"`
	ShippingStatus string             `json:"shipping_status" bson:"shipping_status"`
}
