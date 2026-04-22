package repositories

import (
	"context"

	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/config"
	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type CategoryRepository struct{}

func (r CategoryRepository) FindAll(ctx context.Context) ([]entity.Category, error) {
	var cats []entity.Category

	cursor, err := config.DB.Collection("categories").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &cats); err != nil {
		return nil, err
	}

	return cats, nil
}