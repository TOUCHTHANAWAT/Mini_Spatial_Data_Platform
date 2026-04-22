package services

import (
	"context"
	"errors"

	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/entity"
	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

type PlaceService struct {
	Repo repositories.PlaceRepository
}

func (s *PlaceService) Create(ctx context.Context, place entity.Place) error {
	if place.Name == "" {
		return errors.New("name is required")
	}
	return s.Repo.Create(ctx, place)
}

func (s *PlaceService) GetAll(ctx context.Context, search string, category string) ([]entity.Place, error) {
	return s.Repo.FindAll(ctx, search, category)
}

func (s *PlaceService) Delete(ctx context.Context, id string) error {
	return s.Repo.Delete(ctx, id)
}

func (s *PlaceService) Update(ctx context.Context, id string, data bson.M) error {
	return s.Repo.Update(ctx, id, data)
}

func (s *PlaceService) FindWithin(ctx context.Context, polygon interface{}) ([]entity.Place, error) {
	return s.Repo.FindWithin(ctx, polygon)
}