package services

import (
	"Mongo-GoClient/database"
	"Mongo-GoClient/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetSingleCustomerService(str string) ([]models.Customer, error) {
	ctx := context.Background()
	customer_cols, err := database.GetCollection(CUSTOMERS_COLLECTION)
	if err != nil {
		fmt.Println("error getting customer collection")
		return nil, fmt.Errorf("unable to fetch customer collection")
	}

	customer_obj_id, err := primitive.ObjectIDFromHex(str)
	// fmt.Println(customer_obj_id)
	if err != nil {
		fmt.Println("can not create customer object id from string")
		return nil, fmt.Errorf("not a valid customer id")
	}

	cursor, err := customer_cols.Find(ctx, bson.D{{"_id", customer_obj_id}})
	// fmt.Println("cursor: ", cursor)

	if err != nil {
		fmt.Println("error finding customer id , creating cursor")
		return nil, fmt.Errorf("no valid customer found")
	}
	var customer_details models.Customer
	for cursor.Next(ctx) {
		if err := cursor.Decode(&customer_details); err != nil {
			fmt.Println("error decoding customer details")
			return nil, fmt.Errorf("can not fetch customer details")
		}
	}
	fmt.Println(customer_details)
	return []models.Customer{customer_details}, nil
}
