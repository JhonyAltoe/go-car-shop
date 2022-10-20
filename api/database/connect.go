package connect

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewCollection(databaseName, collectionName string) *mongo.Collection {
	/*
		Connect to my cluster
	*/
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:password@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	/*
		List databases
	*/
	collection := client.Database(databaseName).Collection(collectionName)
	return collection
}
