package database

import "go.mongodb.org/mongo-driver/mongo"

const Users = "users"

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}
