package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongo_url string = "mongodb://localhost:27017/"

func connectToMongo() {

	// MongoDB Connection URI
	clientOptions := options.Client().ApplyURI(mongo_url)

	// Set 	a timeout using context
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	// Connectto MongoDB and Ping the Server
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := client.Database("api").Collection("users")

	var result bson.M
	filter := bson.M{"name": "alex"}
	err = collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found document with id: ", result["_id"])
}
