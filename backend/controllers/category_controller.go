package controllers

import (
	// "net/http"

	"github.com/labstack/echo/v4"
	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/services"
)

type CategoryController struct {
	Service services.CategoryService
}

func (cc *CategoryController) GetAll(c echo.Context) error {
	cats, err := cc.Service.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(500, map[string]string{
			"error": "failed to fetch categories",
		})
	}

	return c.JSON(200, cats)
}