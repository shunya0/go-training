package main

import (
	"Mongo-GoClient/database"
	"Mongo-GoClient/handlers"
	"context"
	"fmt"
	"net/http"
)

func main() {
	defer database.Client.Disconnect(context.Background())
	port := ":8080"
	http.HandleFunc("/orders", handlers.GetOrderById)

	fmt.Println("Starting Server at", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		panic(fmt.Errorf("Error starting server: ", err))

	}

}
