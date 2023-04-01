package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Database = "ci_watcher"
const uri = "mongodb://192.168.50.244:27017"

var client *mongo.Client

func init() {
	// use v1 API
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	c, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	client = c
}

func GetClient() *mongo.Client {
	return client
}

func GetCollection(name string) *mongo.Collection {
	return client.Database(Database).Collection(name)
}

var Owner = GetCollection("owner")
var Repo = GetCollection("repo")
var Commit = GetCollection("commit")
var Version = GetCollection("version")
var Task = GetCollection("task")
