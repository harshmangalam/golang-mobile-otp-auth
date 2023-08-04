package util

import (
	"auth/config"
	"auth/database"
	"auth/model"
	"auth/schema"
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindUserByPhone(phone string) (*model.User, error) {
	// Create a context and a collection instance
	ctx := context.TODO()
	collection := database.Mg.Db.Collection(database.Users)

	// Create a filter to find the user by phone number
	filter := bson.M{"phone": phone}

	// Create a variable to store the result
	var result model.User

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

func UpdateUser(userID primitive.ObjectID, updatedFields map[string]any) error {
	// Create a context and a collection instance
	ctx := context.TODO()
	collection := database.Mg.Db.Collection(database.Users)

	// Create a filter to find the user by ID
	filter := bson.M{"_id": userID}

	// Create an update with the provided fields
	update := bson.M{"$set": updatedFields}

	// Update the user document in the collection
	_, err := collection.UpdateOne(ctx, filter, update)
	return err

}

func FindUserById(userId string) (*model.User, error) {
	// Create a context and a collection instance
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	ctx := context.TODO()
	collection := database.Mg.Db.Collection(database.Users)

	// Create a filter to find the user by phone number
	filter := bson.M{"_id": id}

	// Create a variable to store the result
	var result model.User

	// Find the user with the given phone number
	err = collection.FindOne(ctx, filter).Decode(&result)

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

func GenerateRandomNumber() string {
	// Generate a random number between 1000 and 9999 (inclusive)
	num := rand.Intn(9000) + 1000
	return strconv.Itoa(num)
}

func GenerateJWT(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	return token.SignedString([]byte(config.Config("SECRET")))

}
