package controllers

import (
	"net/http"

	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/entity"
	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/services"
	// "github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/utils"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlaceController struct {
	Service services.PlaceService
}

func toFeature(p entity.Place) entity.Feature {
	return entity.Feature{
		Type: "Feature",
		ID:   p.ID.Hex(),
		Geometry: entity.Geometry{
			Type:        p.Location.Type,
			Coordinates: p.Location.Coordinates,
		},
		Properties: map[string]interface{}{
			"name": p.Name,
		},
	}
}

func validatePlace(p entity.Place) string {
	if p.Name == "" {
		return "name is required"
	}
	if p.Location.Type != "Point" {
		return "location type must be Point"
	}

	lng := p.Location.Coordinates[0]
	lat := p.Location.Coordinates[1]

	if lng < -180 || lng > 180 {
		return "invalid longitude"
	}
	if lat < -90 || lat > 90 {
		return "invalid latitude"
	}

	return ""
}

type FeatureRequest struct {
	Type       string `json:"type"`
	Geometry   entity.Geometry `json:"geometry"`
	Properties struct {
		Name string `json:"name"`
	} `json:"properties"`
}

func (pc *PlaceController) CreateFeature(c echo.Context) error {
	var req FeatureRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request format",
		})
	}

	place := entity.Place{
		ID:       primitive.NewObjectID(),
		Name:     req.Properties.Name,
		Location: req.Geometry,
	}

	if msg := validatePlace(place); msg != "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": msg,
		})
	}

	if err := pc.Service.Create(c.Request().Context(), place); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to create place",
		})
	}

	return c.JSON(http.StatusCreated, toFeature(place))
}

func (pc *PlaceController) GetFeatures(c echo.Context) error {
	places, err := pc.Service.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch places",
		})
	}

	var features []entity.Feature
	for _, p := range places {
		features = append(features, toFeature(p))
	}

	return c.JSON(http.StatusOK, entity.FeatureCollection{
		Type:     "FeatureCollection",
		Features: features,
	})
}

func (pc *PlaceController) Collections(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"collections": []map[string]string{
			{
				"id":    "places",
				"title": "Places",
			},
		},
	})
}