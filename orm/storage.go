package orm

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v3"
)

const Database = "ci_watcher"
const uri = "mongodb://localhost:27017"

var client *mongo.Client
var OwnerColl *mongo.Collection
var CaseColl *mongo.Collection

func init() {
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := c.Database("admin").RunCommand(
		context.TODO(),
		bson.D{
			{Key: "ping", Value: 1},
		},
	).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to MongoDB!")

	client = c
	OwnerColl = GetCollection("owner")
	CaseColl = GetCollection("case")
}

func GetCollection(name string) *mongo.Collection {
	return client.Database(Database).Collection(name)
}

func Migrate() {
	var Schd Schd
	f, err := os.ReadFile("./cases/schd.yaml")
	if err != nil {
		fmt.Println("读取文件失败：", err)
		panic(err)
	}

	yaml.Unmarshal(f, &Schd)

	_, err = OwnerColl.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	for _, owner := range Schd.Owners {
		// fmt.Println(owner)
		OwnerColl.InsertOne(context.Background(), owner)
	}

	_, err = CaseColl.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	for _, caseInfo := range Schd.Cases {
		// fmt.Println(caseInfo)
		caseInfo.Status = "pending"
		caseInfo.Result = ""
		client.Database(Database).Collection("case").InsertOne(context.Background(), caseInfo)
	}
}
