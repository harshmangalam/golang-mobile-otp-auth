package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mg MongoInstance

func Connect() error {
	dbName := "mobile_auth"
	uri := "mongodb://127.0.0.1:27017/" + dbName
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	db := client.Database(dbName)
	if err != nil {
		return err
	}
	Mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}
