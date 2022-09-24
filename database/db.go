package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
	users  *mongo.Collection
}

// make env vars
var mongoUri string = "mongodb+srv://admin:Lb0eMDLh0g6yY2Yp@cluster0.5zlwicm.mongodb.net/?retryWrites=true&w=majority"
var mongoDb string = "hackDFW"

func ConnectDB() (*DB, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Print("\nDB connection failed in database package")
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Connect(ctx)
	return &DB{
		client: client,
		users:  client.Database(mongoDb).Collection("users"),
	}, nil
}
