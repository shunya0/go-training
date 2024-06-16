package services

import (
	"Mongo-GoClient/database"
	"Mongo-GoClient/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBill(customer_id primitive.ObjectID, bill models.BillGen) ([]primitive.ObjectID, error) {
	ctx := context.Background()

	bill_cols, err := database.GetCollection(BILL_COLLECTION)
	if err != nil {
		fmt.Println("error getting bill collection", err)
		return nil, fmt.Errorf("could not get bill collection")
	}

	cursor, err := bill_cols.InsertOne(ctx, bill)
	if err != nil {
		fmt.Println("error creating bill")
		return nil, fmt.Errorf("bill not created")
	}

	return []primitive.ObjectID{cursor.InsertedID.(primitive.ObjectID)}, nil
}
