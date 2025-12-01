package services

import (
	"context"
	"db-connection/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// InsertUser inserts a new user into the database.
func InsertUser(client *mongo.Client, user models.User) (*mongo.InsertOneResult, error) {
    collection := client.Database("testdb").Collection("users")
    return collection.InsertOne(context.TODO(), user)
}

// FindUsers queries users from the database.
func FindUsers(client *mongo.Client, filter interface{}) ([]models.User, error) {
    collection := client.Database("testdb").Collection("users")
    cursor, err := collection.Find(context.TODO(), filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    var users []models.User
    if err := cursor.All(context.TODO(), &users); err != nil {
        return nil, err
    }
    return users, nil
}