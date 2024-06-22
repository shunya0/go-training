package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserAddress struct {
	City string `json:"city" bson:"city"`
	Zip  string `json:"zip" bson:"zip"`
}

type ItemOrdered struct {
	ProductId  primitive.ObjectID `json:"product_id" bson:"product_id"`
	Quantity   int                `json:"quantity" bson:"quantity"`
	Price      int                `json:"price" bson:"price"`
	DiscountId primitive.ObjectID `json:"discount_id" bson:"discount_id"`
}

type BillGen struct {
	OrderID        primitive.ObjectID `json:"order_id" bson:"order_id"`
	CustomerID     primitive.ObjectID `json:"customer_id" bson:"customer_id"`
	ShippingStatus string             `json:"shipping_status" bson:"shipping_status"`
	Address        UserAddress        `json:"address" bson:"address"`
	Items          []ItemOrdered      `json:"orders" bson:"orders"`
}

//used for controllers

type CreateBill struct {
	CustomerId string               `json:"customer_id" bson:"customer_id" binding:"required,customer_id"`
	Products   []ProductDetailsBill `json:"products"`
}
type ProductDetailsBill struct {
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type GetBill struct {
	BillId string `json:"bill_id" bson:"bill_id" binding:"required,bill_id"`
}
type GetBillByCustomer struct {
	CustomerID string `json:"customer_id" bson:"customer_id" binding:"required,customer_id"`
}
