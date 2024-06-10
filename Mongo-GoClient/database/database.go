package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func init() {
	const URI string = "mongodb://127.0.0.1:27017/"

	ctx := context.Background()
	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		panic(fmt.Errorf("Error while connecting to mongoDb (database.go): ", err))

	}
	
	if err = Client.Ping(ctx, nil); err != nil {
		panic(fmt.Errorf("Client not pingged (database.go): ", err))
	}

	fmt.Println("Connected To MongoDb!")

}

func GetCollection(collectionName string) (*mongo.Collection, error) {
	if Client == nil {
		return nil, fmt.Errorf("Client Not connected (database.go)")

	}

	return Client.Database("tasksheet").Collection(collectionName), nil

}
