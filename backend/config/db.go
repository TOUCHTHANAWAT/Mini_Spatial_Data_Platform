package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func InitDB() {
	// 👇 ใช้ localhost ตามที่คุณต้องการ
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

	log.Println("MongoDB connected (localhost) 🚀")
}