package repositories

import (
	"context"

	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/config"
	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/entity"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type PlaceRepository struct{}

func (r PlaceRepository) Create(ctx context.Context, place entity.Place) error {
	_, err := config.DB.Collection("places").InsertOne(ctx, place)
	return err
}

func (r PlaceRepository) FindAll(ctx context.Context, search string, category string) ([]entity.Place, error) {
	var places []entity.Place

	filter := bson.M{
		"deletedAt": bson.M{
			"$exists": false,
		},
	}

	//search filter
	if search != "" {
		filter["name"] = bson.M{
			"$regex":   search,
			"$options": "i",
		}
	}

	// category filter (ใส่ตรงนี้)
	if category != "" {
		filter["category"] = category
	}

	cursor, err := config.DB.Collection("places").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &places); err != nil {
		return nil, err
	}

	return places, nil
}

func (r PlaceRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = config.DB.Collection("places").UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{
			"$set": bson.M{
				"deletedAt": time.Now(),
			},
		},
	)

	return err
}

func (r PlaceRepository) Update(ctx context.Context, id string, data bson.M) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = config.DB.Collection("places").UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{
			"$set": data,
		},
	)

	return err
}

func (r PlaceRepository) FindWithin(ctx context.Context, polygon interface{}) ([]entity.Place, error) {
	var places []entity.Place

	filter := bson.M{
		"location": bson.M{
			"$geoWithin": bson.M{
				"$geometry": polygon,
			},
		},
		"deletedAt": bson.M{
			"$exists": false,
		},
	}

	cursor, err := config.DB.Collection("places").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	cursor.All(ctx, &places)

	return places, nil
}