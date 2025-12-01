package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToMongoDB connects to MongoDB and returns a client.
func ConnectToMongoDB() (*mongo.Client, error) {

	// Set client options
	clientOptions :=options.Client().ApplyURI("mongodb://localhost:27017")

	// connect to MongoDB

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	fmt.Println("Connecting to MongoDB")

	return client, nil

}

// CloseMongoDB disconnects from MongoDB.

func CloseMongoDB(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Disconnected from MongoDB.")
}