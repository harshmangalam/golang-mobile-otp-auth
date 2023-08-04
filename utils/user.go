package utils

import (
	"auth/database"
	"auth/models"
	"auth/schema"
	"context"
	"math/rand"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindUserByPhone(phone string) (*models.User, error) {
	// Create a context and a collection instance
	ctx := context.TODO()
	collection := database.Mg.Db.Collection(database.Users)

	// Create a filter to find the user by phone number
	filter := bson.M{"phone": phone}

	// Create a variable to store the result
	var result models.User

	// Find the user with the given phone number
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// If the error is ErrNoDocuments, it means no user was found
			return nil, nil
		}
		// Handle other potential errors
		return nil, err
	}

	return &result, nil
}

func InsertUser(user *schema.RegisterBody) (any, error) {
	// Create a context and a collection instance
	ctx := context.TODO()
	collection := database.Mg.Db.Collection(database.Users)

	// Insert the user into the collection
	result, err := collection.InsertOne(ctx, user)
	return result.InsertedID, err
}

func GenerateRandomNumber() int {
	// Generate a random number between 1000 and 9999 (inclusive)
	return rand.Intn(9000) + 1000
}
