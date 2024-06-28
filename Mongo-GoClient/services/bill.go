package services

import (
	"Mongo-GoClient/database"
	"Mongo-GoClient/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
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

func GetBill(bill_id primitive.ObjectID) ([]models.BillGen, error) {
	ctx := context.Background()

	bill_col, err := database.GetCollection(BILL_COLLECTION)
	if err != nil {
		fmt.Println("error getting bill collection", err)
		return nil, fmt.Errorf("could not get bill collection")
	}

	cursor, err := bill_col.Find(ctx, bson.D{{"_id", bill_id}})
	if err != nil {
		fmt.Println("error finding bill")
		return nil, fmt.Errorf("bill not found")
	}

	var bill_details []models.BillGen
	for cursor.Next(ctx) {
		var bill_detail models.BillGen
		err := cursor.Decode(&bill_detail)
		if err != nil {
			fmt.Println("unable to decode bill detial")
			return nil, fmt.Errorf("unable to show bill details")
		}
		bill_details = append(bill_details, bill_detail)

	}

	return bill_details, nil
}

func GetBillByCustomerId(customer_id primitive.ObjectID) ([]models.BillGen, error) {
	ctx := context.Background()
	bill_col, err := database.GetCollection(BILL_COLLECTION)
	if err != nil {
		fmt.Println("error getting bill collection", err)
		return nil, fmt.Errorf("could not get bill collection")
	}

	cursor, err := bill_col.Find(ctx, bson.D{{"customer_id", customer_id}})
	if err != nil {
		fmt.Println("error finding bill")
		return nil, fmt.Errorf("bill not found")
	}

	var bill_details []models.BillGen
	for cursor.Next(ctx) {
		var bill_detail models.BillGen
		err := cursor.Decode(&bill_detail)
		if err != nil {
			fmt.Println("unable to decode bill detial")
			return nil, fmt.Errorf("unable to show bill details")
		}
		bill_details = append(bill_details, bill_detail)

	}

	return bill_details, nil
}
