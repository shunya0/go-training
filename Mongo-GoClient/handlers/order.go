package handlers

import (
	"Mongo-GoClient/database"
	"Mongo-GoClient/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const ORDER_COLLECTION string = "orders"
const DISCOUNT_COLLECTION string = "discount"
const CUSTOMERS_COLLECTION string = "customers"
const SHIPPING_COLLECTION string = "shipping"
const PRODUCTS_COLLECTION string = "products"

func getOrders(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	col, err := database.GetCollection("orders")
	if err != nil {
		http.Error(w, "failed to get collection (main.go/getOrders) ", http.StatusInternalServerError)
		return
	}
	fmt.Println(database.Client)

	fmt.Println(col)
	cursor, err := col.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Err", err)
		http.Error(w, "Error finding collection (main.go/getOrders)", http.StatusInternalServerError)

		return
	}

	defer cursor.Close(ctx)

	var orders []models.Order
	for cursor.Next(ctx) {
		var order models.Order
		err := cursor.Decode(&order)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error decoding order (main.go/getOrders) ", http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)

	}

	if err := cursor.Err(); err != nil {
		http.Error(w, "Error in iterating cursor (main.go/getOrders) ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		http.Error(w, "Error encoding orders (main.go/getOrders)", http.StatusInternalServerError)
		return

	}
}
func GetOrderById(w http.ResponseWriter, r *http.Request) {
	order_id_str := r.URL.Query().Get("orderid")
	order_id, err := primitive.ObjectIDFromHex(order_id_str)
	if err != nil {
		http.Error(w, "Error converting string to objectId (order.go)", http.StatusInternalServerError)
		return
	}

	ctx := context.Background()

	ordersCol, err := database.GetCollection(ORDER_COLLECTION)
	if err != nil {
		http.Error(w, "Error getting collection: (order.go/orders)", http.StatusInternalServerError)
		return
	}

	pipeline := mongo.Pipeline{

		{
			bson.E{"$match", bson.D{{"_id", order_id}}},
		},
		{
			bson.E{"$lookup", bson.D{{"from", "customers"}, {"localField", "customer_id"}, {"foreignField", "_id"}, {"as", "customer_details"}}},
		},
		{
			bson.E{"$unwind", "$customer_details"},
		},
		{
			bson.E{"$unwind", "$items"},
		},
		{
			bson.E{"$lookup", bson.D{{"from", "products"}, {"localField", "items.product_id"}, {"foreignField", "_id"}, {"as", "product_details"}}},
		},
		{
			bson.E{"$unwind", "$product_details"},
		},
		{
			bson.E{"$lookup", bson.D{{"from", "discount"}, {"localField", "items.discount_id"}, {"foreignField", "_id"}, {"as", "discount_details"}}},
		},
		{
			bson.E{"$unwind", "$discount_details"},
		},
		{
			bson.E{"$group", bson.D{{"_id", "$_id"}, {"customer_name", bson.D{{"$first", "$customer_details.name"}}},
				{"customer_email", bson.D{{"$first", "$customer_details.email"}}}, {"join_date", bson.D{{"$first", "$customer_details.join_date"}}},
				{"address", bson.D{{"$first", "$customer_details.address"}}}, {"ordered_date", bson.D{{"$first", "$order_date"}}}, {"status", bson.D{{"$first", "$status"}}},
				{"order_details", bson.D{{"$push", bson.D{{"name", "$product_details.name"}, {"quantity", "$items.quantity"}, {"original_price_per_product", "$items.price"},
					{"discount", bson.D{{"$multiply", bson.A{"$items.price", bson.D{{"$divide", bson.A{"$discount_details.percentage", 100}}}}}}}}}}}}},
		},
	}

	cursor, err := ordersCol.Aggregate(ctx, pipeline)
	if err != nil {
		fmt.Fprintln(w, err)
		http.Error(w, "Error aggreagating orders: (order.go/cursor)", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var orders []models.OrderAggergate

	for cursor.Next(ctx) {
		var order models.OrderAggergate
		errS := cursor.Decode(&order)
		if errS != nil {
			http.Error(w, "Error decoding cursor: (order.go/cursor.next)", http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)
	}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(orders); err != nil {
		http.Error(w, "Error encoding orders ,(order.go/encode)", http.StatusInternalServerError)
		return
	}

}
