package main

import (
	"Mongo-GoClient/database"
	"Mongo-GoClient/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func getOrders(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	col, err := database.GetCollection("orders")
	if err != nil {
		http.Error(w, "failed to get collection (main.go): ", http.StatusInternalServerError)
		return
	}
	fmt.Println(database.Client)

	fmt.Println(col)
	cursor, err := col.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Err", err)
		http.Error(w, "Error finding collection (main.go): ", http.StatusInternalServerError)

		return
	}

	defer cursor.Close(ctx)

	var orders []models.Order
	for cursor.Next(ctx) {
		var order models.Order
		err := cursor.Decode(&order)
		if err != nil {
			http.Error(w, "Error decoding order (main.go): ", http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)

	}

	if err := cursor.Err(); err != nil {
		http.Error(w, "Error in iterating cursor (main.go): ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		http.Error(w, "Error encoding orders (main.go): ", http.StatusInternalServerError)
		return

	}
}

func main() {
	defer database.Client.Disconnect(context.Background())
	port := ":8080"

	http.HandleFunc("/orders", getOrders)

	fmt.Println("Starting Server at", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		panic(fmt.Errorf("Error starting server: ", err))

	}

}
