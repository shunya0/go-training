package services

import (
	"Mongo-GoClient/database"
	"Mongo-GoClient/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOrderService(order models.Order) ([]primitive.ObjectID, error) {
	ctx := context.Background()

	order_cols, err := database.GetCollection(ORDER_COLLECTION)
	if err != nil {
		fmt.Println("error getting order collection")
		return nil, fmt.Errorf("error fetching collection")
	}

	cursor, err := order_cols.InsertOne(ctx, order)
	if err != nil {
		fmt.Println("can not create order")
		return nil, fmt.Errorf("order not created")
	}
	return []primitive.ObjectID{cursor.InsertedID.(primitive.ObjectID)}, nil
}
