package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// deals with models
type Items struct {
	ProductId  primitive.ObjectID `json:"product_id" bson:"product_id"`
	Quantity   int                `json:"quantity" bson:"quantity"`
	Price      int                `json:"price" bson:"price"`
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

// deals with request
type CreateOrderProductBody struct {
	ProductId string `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}

type CreateOrderRequestBody struct {
	CustomerId string                   `json:"customer_id" binding:"required,customer_id"`
	Products   []CreateOrderProductBody `json:"products" binding:"required,product_id"`
}

type CancelOrderBody struct {
	CustomerId string `json:"customer_id" binding:"required,customer_id"`
	OrderId    string `json:"order_id" binding:"required,order_id"`
	ShippingId string `json:"shipping_id" binding:"required,shipping_id"`
}
