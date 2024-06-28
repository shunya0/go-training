package services

import (
	"Mongo-GoClient/database"
	"Mongo-GoClient/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func GetDiscountServices() ([]models.Discount, error) {
	ctx := context.Background()

	discountCols, err := database.GetCollection(DISCOUNT_COLLECTION)
	if err != nil {
		fmt.Println("error getting discountCols", err)
		return nil, fmt.Errorf("unable to get discounts")
	}

	cursor, err := discountCols.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("error creating cursor for discountCols", err)

		return nil, fmt.Errorf("error in creating cursor")
	}
	defer cursor.Close(ctx)

	var discountArr []models.Discount
	for cursor.Next(ctx) {
		var discount models.Discount
		if err := cursor.Decode(&discount); err != nil {
			fmt.Println("error decodind discount", err)

			return nil, fmt.Errorf("unable to decode collection")
		}
		discountArr = append(discountArr, discount)
	}

	return discountArr, nil

}
