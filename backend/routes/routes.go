package routes

import (
	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/controllers"
	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/repositories"
	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/services"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	placeRepo := repositories.PlaceRepository{}
	placeService := services.PlaceService{Repo: placeRepo}
	placeController := controllers.PlaceController{Service: placeService}

	api := e.Group("/api")

	api.GET("/collections", placeController.Collections)
	api.GET("/collections/places/items", placeController.GetFeatures)
	api.POST("/collections/places/items", placeController.CreateFeature)
	api.PATCH("/collections/places/items/:id", placeController.UpdateFeature)
	api.DELETE("/collections/places/items/:id", placeController.DeleteFeature)

	api.POST("/collections/places/within", placeController.FindWithin)

	categoryRepo := repositories.CategoryRepository{}
	categoryService := services.CategoryService{Repo: categoryRepo}
	categoryController := controllers.CategoryController{Service: categoryService}

	api.GET("/categories", categoryController.GetAll)
}
