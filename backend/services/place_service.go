package services

import (
	"context"
	"errors"

	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/entity"
	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/repositories"
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

func (s PlaceService) GetAll(ctx context.Context) ([]entity.Place, error) {
	return s.Repo.FindAll(ctx)
}