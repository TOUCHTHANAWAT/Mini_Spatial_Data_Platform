package seed

import (
	"context"
	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/config"
	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/entity"
)

func SeedCategories() error {
	col := config.DB.Collection("categories")

	count, _ := col.CountDocuments(context.TODO(), map[string]interface{}{})
	if count > 0 {
		return nil // already seeded
	}

	data := []interface{}{
		entity.Category{Value: "hospital", Label: "hospital", Color: "pink"},
		entity.Category{Value: "restaurant", Label: "Restaurant", Color: "red"},
		entity.Category{Value: "cafe", Label: "Cafe",  Color: "brown"},
		entity.Category{Value: "school", Label: "School", Color: "blue"},
	}

	_, err := col.InsertMany(context.TODO(), data)
	return err
}