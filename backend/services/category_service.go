package services

import (
	"context"

	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/entity"
	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/repositories"
)

type CategoryService struct {
	Repo repositories.CategoryRepository
}

func (s CategoryService) GetAll(ctx context.Context) ([]entity.Category, error) {
	return s.Repo.FindAll(ctx)
}