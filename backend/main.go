package main

import (
	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/config"
	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/routes"
	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/seed"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson"

	"context"
	"os"
)

func main() {
	config.InitDB()

	collection := config.DB.Collection("places")

	count, _ := collection.CountDocuments(context.Background(), bson.M{})
	if count == 0 {
		seed.ImportGeoJSON()
		seed.CreateIndex()
	}

	catCollection := config.DB.Collection("categories")

	countCat, _ := catCollection.CountDocuments(context.Background(), bson.M{})
	if countCat == 0 {
		seed.SeedCategories()
	}

	e := echo.New()
	e.Use(middleware.CORS())

	// e.Use(middleware.Logger()) //เก็บและแสดง log
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "\033[32m${status}\033[0m | ${latency_human} | ${method} ${uri}\n",
	}))
	e.Use(middleware.Recover()) //กัน server ล่ม

	routes.InitRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
