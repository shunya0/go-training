package services

import (
	"Mongo-GoClient/database"
	"Mongo-GoClient/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateShippingDetailsService() ([]primitive.ObjectID, error) {
	ctx := context.Background()

	shippingCol, err := database.GetCollection(SHIPPING_COLLECTION)
	if err != nil {
		fmt.Println("error getting shipping collections", err)
		return nil, fmt.Errorf("unable to fetch shippingCol")
	}

	var shipping_doc models.ShippingCreation
	shipping_doc.Status = "pending"
	cursor, err := shippingCol.InsertOne(ctx, shipping_doc)
	if err != nil {
		fmt.Println("error in creating cursor for shippingCol", err)
		return nil, fmt.Errorf("can not create cursor for shipping")
	}
	return []primitive.ObjectID{cursor.InsertedID.(primitive.ObjectID)}, nil
}

func UpdateShippingOrderIdService(shipping_id primitive.ObjectID, order_id primitive.ObjectID) error {
	ctx := context.Background()

	shipping_cols, err := database.GetCollection(SHIPPING_COLLECTION)
	if err != nil {
		fmt.Println("can not get shipping collection", err)
		return fmt.Errorf("unable to fetch shipping collection ")
	}

	filter := bson.D{{"_id", shipping_id}}
	cursor, err := shipping_cols.UpdateOne(ctx, filter, bson.D{{"$set", bson.D{{"order_id", order_id}}}})
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

func DeleteShippmentService(shipping_id primitive.ObjectID) error {

	ctx := context.Background()

	shipping_col, err := database.GetCollection(SHIPPING_COLLECTION)
	if err != nil {
		fmt.Println("can not get shipping collection:(deleteShippmentService) ", err)
		return fmt.Errorf("unable to fetch shipping collection: (deleteShippmentService) ")
	}

	cursor, err := shipping_col.DeleteOne(ctx, bson.D{{"_id", shipping_id}})
	if err != nil {
		fmt.Println("error deleting order id")
		return fmt.Errorf("shipping id not deleted")

	}
	if cursor.DeletedCount == 0 {
		fmt.Println("invalid user")
		return fmt.Errorf("shipping id not deleted")
	}
	return nil
}

func GetShippingDetailsService(shipping_id primitive.ObjectID) ([]models.Shipping, error) {
	ctx := context.Background()

	shipping_col, err := database.GetCollection(SHIPPING_COLLECTION)
	if err != nil {
		fmt.Println("can not get shipping collection:(GetShippingDetailsService) ", err)
		return nil, fmt.Errorf("unable to fetch shipping colleciton (GetShippingDetailsService)")
	}

	cursor, err := shipping_col.Find(ctx, bson.D{{"_id", shipping_id}})
	if err != nil {
		fmt.Println("can not find shipping services for this collection (GetShippingDetailsService)")
		return nil, fmt.Errorf("shipping collection not found (GetShippingDetailsService)")
	}

	var shipping_details models.Shipping
	for cursor.Next(ctx) {
		err := cursor.Decode(&shipping_details)
		if err != nil {
			fmt.Println("unable to decode shipping detials(GetShippingDetailsService)")
			return nil, fmt.Errorf("unable to show shipping details (GetShippingDetailsService)")
		}
	}

	return []models.Shipping{shipping_details}, nil
}
