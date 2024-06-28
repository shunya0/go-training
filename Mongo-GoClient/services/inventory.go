package services

import (
	"Mongo-GoClient/database"
	"Mongo-GoClient/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateInventoryService(products_arr []models.Items) error {
	ctx := context.Background()

	invCols, err := database.GetCollection(INVENTORY_COLLECTION)
	if err != nil {
		fmt.Println("error getting inv", err)
		return fmt.Errorf("unable to get inventory")
	}

	for _, product := range products_arr {

		cursor_find, err := invCols.Find(ctx, bson.D{{"product_id", product.ProductId}})
		if err != nil {
			fmt.Println("error setting up cursor", err)
			return fmt.Errorf("error in creating cursor")
		}

		var fetch_inv_doc models.Inventory
		for cursor_find.Next(ctx) {
			if err := cursor_find.Decode(&fetch_inv_doc); err != nil {
				fmt.Println("error decoding cursor for inv_doc", err)
				return fmt.Errorf("error while decoding inventory")
			}
		}
		fmt.Println(fetch_inv_doc)
		fmt.Println(product)
		if fetch_inv_doc.QuantityAvailable < int(product.Quantity) {
			fmt.Println("if")
			fetch_inv_doc.QuantityAvailable = 0
		} else {
			fmt.Println("else")
			fetch_inv_doc.QuantityAvailable -= int(product.Quantity)
		}

		filter := bson.D{{"product_id", product.ProductId}}
		cursor_update, err := invCols.UpdateOne(ctx, filter, bson.D{{"$set", bson.D{{"quantity_available", fetch_inv_doc.QuantityAvailable}}}})
		if err != nil {
			fmt.Println("error while updating inventory,", err)
			return fmt.Errorf("can not update inventory")
		}
		fmt.Println(cursor_update)
		if cursor_update.MatchedCount != 1 {
			fmt.Println("inventory not updated")
			return fmt.Errorf("inventory update failed")
		}

	}
	return nil
}
