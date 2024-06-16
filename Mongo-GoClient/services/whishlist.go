package services

import (
	"Mongo-GoClient/database"
	"Mongo-GoClient/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetWishlistService(customer_obj_id primitive.ObjectID) ([]models.Whishlist, error) {
	ctx := context.Background()

	whishCol, err := database.GetCollection(WHISHLIST_COLLECTIONS)
	if err != nil {
		return nil, fmt.Errorf("error geting collections")
	}

	cursor, err := whishCol.Find(ctx, bson.D{{"customer_id", customer_obj_id}})
	if err != nil {
		return nil, fmt.Errorf("error fetching Wishlist")
	}

	var wishlists []models.Whishlist
	for cursor.Next(ctx) {
		var wishlist models.Whishlist

		err := cursor.Decode(&wishlist)
		if err != nil {
			return nil, fmt.Errorf("error decoding cursor")
		}

		wishlists = append(wishlists, wishlist)
	}

	return wishlists, nil
}
