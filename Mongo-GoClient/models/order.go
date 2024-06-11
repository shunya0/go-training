package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Items struct {
	ProductId  string             `json:"product_id" bson:"product_id"`
	Quantity   int64              `json:"quantity" bson:"quantity"`
	Price      int64              `json:"price" bson:"price"`
	DiscountId primitive.ObjectID `json:"discount_id" bson:"discount_id"`
}
type Order struct {
	ID         string             `json:"_id,omitempty"  bson:"_id,omitempty"`
	CustomerId primitive.ObjectID `json:"customer_id" bson:"customer_id"`
	OrderDate  primitive.DateTime `json:"order_date" bson:"order_date"`
	Status     string             `json:"status" bson:"status"`
	Items      []Items            `json:"items" bson:"items"`
	ShippingID primitive.ObjectID `json:"shipping_id" bson:"shipping_id"`
}
