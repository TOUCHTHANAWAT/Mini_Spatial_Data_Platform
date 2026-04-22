package seed

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GeoJSON struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string     `json:"type"`
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
}

type Geometry struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type Properties struct {
	Name string `json:"name"`
}

type Place struct {
	Name     string   `bson:"name"`
	Category string   `bson:"category"`
	Location Geometry `bson:"location"`
}


func ImportGeoJSON() {
	file, err := os.Open("data.geojson") // ไฟล์ต้องอยู่ root project
	if err != nil {
		log.Fatal("open file error:", err)
	}
	defer file.Close()

	var geo GeoJSON
	if err := json.NewDecoder(file).Decode(&geo); err != nil {
		log.Fatal("decode error:", err)
	}

	collection := config.DB.Collection("places")

	var docs []interface{}

	for _, f := range geo.Features {
		// กันพัง: ต้องมี name + geometry
		if f.Properties.Name == "" || len(f.Geometry.Coordinates) != 2 {
			continue
		}

		doc := Place{
			Name:     f.Properties.Name,
			Category: "hospital",
			Location: f.Geometry,
		}

		docs = append(docs, doc)
	}

	if len(docs) == 0 {
		log.Println("no valid data to insert")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.InsertMany(ctx, docs)
	if err != nil {
		log.Fatal("insert error:", err)
	}

	log.Println("Import success:", len(docs))
}

func CreateIndex() {
	collection := config.DB.Collection("places")

	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "location", Value: "2dsphere"},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatal("create index error:", err)
	}

	log.Println("Index created (2dsphere)")
}