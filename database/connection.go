package database

import (
	"auth/config"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mg MongoInstance

func Connect() error {
	dbName := config.Config("DATABASE_NAME")
	uri := config.Config("DATABASE_URI") + dbName
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	log.Printf("Connected with databse %s", dbName)
	db := client.Database(dbName)
	Mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}
