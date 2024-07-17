package services

import (
	"Mongo-GoClient/database"
	"Mongo-GoClient/models"
	"Mongo-GoClient/utils"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user models.RegisterUser) ([]primitive.ObjectID, error) {
	ctx := context.Background()
	var existingUser models.RegisterUser

	user_cols, err := database.GetCollection(utils.USER_COLLECTION)
	if err != nil {
		fmt.Println("user_cols: ", err)
		return nil, fmt.Errorf("cursor,user_cols>user.go: ", err)
	}

	filter := bson.M{"username": user.Username}
	err = user_cols.FindOne(ctx, filter).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {

		} else {
			fmt.Println("error checking if user exists: ", err)
			return nil, fmt.Errorf("error checking if user exists: ", err)
		}
	} else {
		return nil, fmt.Errorf("user with username %s and mail %s already exists", user.Username, user.Email)
	}
	filter_mail := bson.M{"email": user.Email}
	err = user_cols.FindOne(ctx, filter_mail).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {

		} else {
			fmt.Println("error checking if user exists: ", err)
			return nil, fmt.Errorf("error checking if user exists: ", err)
		}
	} else {
		return nil, fmt.Errorf("user with username %s and mail %s already exists", user.Username, user.Email)
	}

	cursor, err := user_cols.InsertOne(ctx, user)
	if err != nil {
		fmt.Println("cursor,services>user.go: ", err)
		return nil, fmt.Errorf("cursor,services>user.go: ", err)
	}

	return []primitive.ObjectID{cursor.InsertedID.(primitive.ObjectID)}, nil

}

func CheckUserExistsService(user models.LoginUser) (bool, error) {
	ctx := context.Background()
	var existingUser models.LoginUser

	user_cols, err := database.GetCollection(utils.USER_COLLECTION)
	if err != nil {
		fmt.Println("Error in getting user collection")
		return false, fmt.Errorf("Error getting user collection", err)
	}

	filter := bson.D{{"email", user.Email}}
	err = user_cols.FindOne(ctx, filter).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {

		} else {
			fmt.Println("error checking if user exists: ", err)
			return false, fmt.Errorf("error checking if user exists", err)
		}
	} else {

		return true, nil // user found
	}
	return false, nil
}

func UserValid(user models.LoginUser) (bool, error) {
	ctx := context.Background()
	var existing_user models.LoginUser
	user_cols, err := database.GetCollection(utils.USER_COLLECTION)
	if err != nil {
		return false, fmt.Errorf("error getting user collection")
	}

	filter := bson.D{{"email", user.Email}, {"password", user.Password}}

	err = user_cols.FindOne(ctx, filter).Decode(&existing_user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, fmt.Errorf("no user found with this mail/password")
		} else {
			fmt.Println("error checking if user exists: ", err)
			return false, fmt.Errorf("error checking if user exists", err)
		}

	} else {

		return true, nil // user found
	}

}
