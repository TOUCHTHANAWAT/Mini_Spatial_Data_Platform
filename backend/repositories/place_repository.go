package repositories

import (
	"context"

	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/config"
	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type PlaceRepository struct{}

func (r PlaceRepository) Create(ctx context.Context, place entity.Place) error {
	_, err := config.DB.Collection("places").InsertOne(ctx, place)
	return err
}

func (r PlaceRepository) FindAll(ctx context.Context) ([]entity.Place, error) {
	var places []entity.Place

	cursor, err := config.DB.Collection("places").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &places); err != nil {
		return nil, err
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return places, nil
}