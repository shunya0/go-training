package handlers

import (
	"Mongo-GoClient/database"
	"Mongo-GoClient/models"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	categories_str := r.URL.Query().Get("categories")
	var categories_arr []string
	categories_bson_arr := bson.A{}

	if len(categories_str) != 0 {
		categories_arr = strings.Split(categories_str, ",")

	}

	for _, v := range categories_arr {
		categories_bson_arr = append(categories_bson_arr, v)
	}

	ctx := context.Background()

	productCol, err := database.GetCollection(PRODUCTS_COLLECTION)
	if err != nil {
		http.Error(w, "failed to get collection (main.go/getProducts): ", http.StatusInternalServerError)
		return
	}

	pipeline_category := mongo.Pipeline{
		{
			bson.E{"$match", bson.D{{"category", bson.D{{"$in", categories_bson_arr}}}}},
		},
		{
			bson.E{"$lookup", bson.D{{"from", "inventory"}, {"localField", "inventory_id"}, {"foreignField", "_id"}, {"as", "inventory_details"}}},
		},
		{
			bson.E{"$unwind", "$inventory_details"},
		},
		{
			bson.E{"$group", bson.D{{"_id", "$_id"}, {"product_name", bson.D{{"$first", "$name"}}},
				{"product_category", bson.D{{"$first", "$category"}}}, {"product_price", bson.D{{"$first", "$price"}}},
				{"quantity_available", bson.D{{"$first", "$inventory_details.quantity_available"}}}}},
		},
		{
			bson.E{"$project", bson.D{{"_id", 0}, {"product_id", "$_id"},
				{"product_name", 1}, {"product_category", 1}, {"product_price", 1},
				{"quantity_available", 1}}},
		},
	}

	pipeline_all := mongo.Pipeline{
		{
			bson.E{"$lookup", bson.D{{"from", "inventory"}, {"localField", "inventory_id"}, {"foreignField", "_id"}, {"as", "inventory_details"}}},
		},
		{
			bson.E{"$unwind", "$inventory_details"},
		},
		{
			bson.E{"$group", bson.D{{"_id", "$_id"}, {"product_name", bson.D{{"$first", "$name"}}},
				{"product_category", bson.D{{"$first", "$category"}}}, {"product_price", bson.D{{"$first", "$price"}}},
				{"quantity_available", bson.D{{"$first", "$inventory_details.quantity_available"}}}}},
		},
		{
			bson.E{"$project", bson.D{{"_id", 0}, {"product_id", "$_id"},
				{"product_name", 1}, {"product_category", 1}, {"product_price", 1},
				{"quantity_available", 1}}},
		},
	}

	var cursor *mongo.Cursor
	var errCursor error
	if len(categories_arr) == 0 {
		cursor, errCursor = productCol.Aggregate(ctx, pipeline_all)

	} else {
		cursor, errCursor = productCol.Aggregate(ctx, pipeline_category)

	}
	if errCursor != nil {
		http.Error(w, "Error finding collection (main.go/getProducts)", http.StatusInternalServerError)
		return
	}

	defer cursor.Close(ctx)

	var products []models.ProductsAggregate
	for cursor.Next(ctx) {
		var product models.ProductsAggregate
		err := cursor.Decode(&product)
		if err != nil {
			http.Error(w, "Error decoding product (main.go/getProducts)", http.StatusInternalServerError)
			return
		}

		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, "Error in iterating cursor (main.go/getProducts)", http.StatusInternalServerError)
		return
	}

	if products == nil {
		http.Error(w, "No product found for this category", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Error encoding products: handlers/product.go", http.StatusInternalServerError)
		return
	}

}
