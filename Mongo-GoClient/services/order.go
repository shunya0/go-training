package services

import (
	"Mongo-GoClient/database"
	"Mongo-GoClient/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
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

func CancelShippmentService(order_id primitive.ObjectID) error {

	ctx := context.Background()

	order_cols, err := database.GetCollection(ORDER_COLLECTION)
	if err != nil {
		fmt.Println("can not get shipping collection:(CancelShippmentService) ", err)
		return fmt.Errorf("unable to fetch shipping collection: (CancelShippmentService) ")
	}

	filter := bson.D{{"_id", order_id}}
	cursor, err := order_cols.UpdateOne(ctx, filter, bson.D{{"$set", bson.D{{"status", "canceled"}}}})
	if err != nil {
		fmt.Println("error updating order id")
		return fmt.Errorf("order id not updated")

	}
	if cursor.MatchedCount != 1 {
		fmt.Println("update not successful")
		return fmt.Errorf("not updated")
	}
	return nil
}

func GetOrderDetialsService(order_id primitive.ObjectID) ([]models.Order, error) {
	ctx := context.Background()

	order_cols, err := database.GetCollection(ORDER_COLLECTION)
	if err != nil {
		fmt.Println("can not get Order collection:(GetOrderDetialsService) ", err)
		return nil, fmt.Errorf("unable to fetch Order colleciton (GetOrderDetialsService)")
	}

	cursor, err := order_cols.Find(ctx, bson.D{{"_id", order_id}})
	if err != nil {
		fmt.Println("can not find Order services for this collection (GetOrderDetialsService)")
		return nil, fmt.Errorf("Order collection not found (GetOrderDetialsService)")
	}

	var order_details models.Order
	for cursor.Next(ctx) {
		err := cursor.Decode(&order_details)
		if err != nil {
			fmt.Println("unable to decode order detials(GetOrderDetialsService)")
			return nil, fmt.Errorf("unable to show order details (GetOrderDetialsService)")
		}
	}

	return []models.Order{order_details}, nil
}

func GetOrderDetialsByCustomerIDService(customer_id primitive.ObjectID) ([]models.Order, error) {
	ctx := context.Background()

	order_cols, err := database.GetCollection(ORDER_COLLECTION)
	if err != nil {
		fmt.Println("can not get Order collection:(GetOrderDetialsService) ", err)
		return nil, fmt.Errorf("unable to fetch Order colleciton (GetOrderDetialsService)")
	}

	cursor, err := order_cols.Find(ctx, bson.D{{"customer_id", customer_id}})
	if err != nil {
		fmt.Println("can not find Order services for this collection (GetOrderDetialsService)")
		return nil, fmt.Errorf("oder collection not found (GetOrderDetialsService)")
	}

	var order_details models.Order
	for cursor.Next(ctx) {
		err := cursor.Decode(&order_details)
		if err != nil {
			fmt.Println("unable to decode order detials(GetOrderDetialsService)")
			return nil, fmt.Errorf("unable to show order details (GetOrderDetialsService)")
		}
	}

	return []models.Order{order_details}, nil
}
