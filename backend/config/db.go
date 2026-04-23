package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	_ = godotenv.Load()

	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	if uri == "" {
		log.Fatal("MONGO_URI is required")
	}

	if dbName == "" {
		log.Fatal("DB_NAME is required")
	}

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

	DB = client.Database(dbName)
	createIndexes()

	log.Println("MongoDB connected")
}
