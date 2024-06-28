package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductsAggregate struct {
	ProductName       string             `json:"product_name" bson:"product_name"`
	ProductCategory   string             `json:"product_category" bson:"product_category"`
	ProductPrice      int                `json:"product_price" bson:"product_price"`
	QuantityAvailable int                `json:"quantity_available" bson:"quantity_available"`
	ProductID         primitive.ObjectID `json:"product_id" bson:"product_id"`
}
