package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	ctx context.Context
	client *mongo.Client
	connection *mongo.Collection
)

func init()  {
	var err error
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	connection = client.Database("test").Collection("test")
}

func insert() {
	id, err := connection.InsertOne(ctx, bson.M{"name": "jack"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id.InsertedID)
}

func queryMany()  {
	cur, err := connection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		var val bson.M
		cur.Decode(&val)
		fmt.Println(val)
	}
}

func queryOne()  {
	filter := bson.M{"name": "jack"}
	val := connection.FindOne(ctx, filter)
	var result bson.M
	val.Decode(&result)
	fmt.Println(result)
}

func deleteOne() {
	connection.DeleteOne(ctx, bson.M{"name": "jack"})
}


func main() {
	// insert()
	deleteOne()
}
