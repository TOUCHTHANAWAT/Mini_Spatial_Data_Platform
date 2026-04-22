package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

var DB *mongo.Database

func createIndexes() {
	collection := DB.Collection("places")

	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "location", Value: "2dsphere"},
		},
	}

	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatal("failed to create index:", err)
	}

	log.Println("2dsphere index created ")
}

func InitDB() {
	uri := "mongodb://localhost:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("create client error:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("connect error:", err)
	}

	// ping (เช็คว่า MongoDB ใช้งานได้จริง)
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("ping error:", err)
	}

	DB = client.Database("spatial_db")
	createIndexes()

	log.Println("MongoDB connected (localhost) ")
}