package infrastructure

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri string
var ctx = context.TODO()

func Open() (*mongo.Client, error) {

	uri = os.Getenv("URI_MONGO_CONNECT")
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	return client, err
}

func GetMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, *mongo.Client, error) {
	client, err := Open()

	if err != nil {
		return nil, nil, err
	}

	collection := client.Database(DbName).Collection(CollectionName)

	return collection, client, nil
}