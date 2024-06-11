package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CustomerAddress struct {
	City string `json:"city" bson:"city"`
	Zip  string `jsno:"zip" bson:"zip"`
}

type ProductDetails struct {
	ProductName          string `json:"name" bson:"name"`
	ProductQuantity      int    `json:"quantity" bson:"quantity"`
	ProductOriginalPrice int    `json:"original_price_per_product" bson:"original_price_per_product"`
	Discount             int    `json:"discount" bson:"discount"`
}

type OrderAggergate struct {
	ID            string             `json:"_id,omitempty"  bson:"_id,omitempty"`
	CustomerName  string             `json:"customer_name" bson:"customer_name"`
	CustomerEmail string             `json:"customer_email" bson:"customer_email"`
	JoinDate      primitive.DateTime `json:"join_date" bson:"join_date"`
	Address       CustomerAddress    `json:"address" bson:"address"`
	Status        string             `json:"status" bson:"status"`
	Order         []ProductDetails   `json:"order_details" bson:"order_details"`
}
