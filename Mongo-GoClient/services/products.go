package services

import (
	"Mongo-GoClient/database"
	"Mongo-GoClient/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ORDER_COLLECTION string = "orders"
const DISCOUNT_COLLECTION string = "discount"
const CUSTOMERS_COLLECTION string = "customers"
const SHIPPING_COLLECTION string = "shipping"
const PRODUCTS_COLLECTION string = "products"

func GetProductsService(product_id_arr_str []string) ([]models.Product, error) {
	var product_id_arr_obj_id []primitive.ObjectID
	for _, product_id_str := range product_id_arr_str {
		temp, err := primitive.ObjectIDFromHex(product_id_str)
		if err != nil {
			return nil, fmt.Errorf("error converting string to object id")
		}
		product_id_arr_obj_id = append(product_id_arr_obj_id, temp)
	}

	ctx := context.Background()

	productCol, err := database.GetCollection(PRODUCTS_COLLECTION)
	if err != nil {
		return nil, fmt.Errorf("error geting collections")
	}

	product_id_arr_bson := bson.A{}
	for _, product_obj_id := range product_id_arr_obj_id {
		product_id_arr_bson = append(product_id_arr_bson, product_obj_id)
	}

	filter := bson.M{"_id": bson.D{{"$in", product_id_arr_bson}}}
	cursor, err := productCol.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error fetching products")
	}

	var products []models.Product
	for cursor.Next(ctx) {
		var product models.Product

		err := cursor.Decode(&product)
		if err != nil {
			return nil, fmt.Errorf("error decoding cursor")
		}

		products = append(products, product)
	}

	return products, nil
}


//