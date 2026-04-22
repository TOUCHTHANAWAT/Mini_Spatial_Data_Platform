package controllers

import (
	"net/http"

	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/entity"
	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/services"
	// "github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/utils"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
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
			"name":     p.Name,
			"category": p.Category,
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
	Type       string          `json:"type"`
	Geometry   entity.Geometry `json:"geometry"`
	Properties struct {
		Name     string `json:"name"`
		Category string `json:"category"`
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
		Category: req.Properties.Category,
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
	search := c.QueryParam("search")
	category := c.QueryParam("category")

	places, err := pc.Service.GetAll(
		c.Request().Context(),
		search,
		category, // ⭐ ส่งต่อ
	)

	if err != nil {
		return c.JSON(500, map[string]string{
			"error": "failed to fetch places",
		})
	}

	features := make([]entity.Feature, 0)

	for _, p := range places {
		features = append(features, toFeature(p))
	}

	// กัน nil
	if features == nil {
		features = []entity.Feature{}
	}

	return c.JSON(200, entity.FeatureCollection{
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

func (pc *PlaceController) DeleteFeature(c echo.Context) error {
	id := c.Param("id")

	if err := pc.Service.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to delete place",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "deleted successfully",
	})
}

type UpdateFeatureRequest struct {
	Geometry   *entity.Geometry `json:"geometry,omitempty"`
	Properties *struct {
		Name     *string `json:"name,omitempty"`
		Category *string `json:"category,omitempty"`
	} `json:"properties,omitempty"`
}

func (pc *PlaceController) UpdateFeature(c echo.Context) error {
	id := c.Param("id")

	var req UpdateFeatureRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": "invalid request"})
	}

	updateData := bson.M{}

	// update name
	if req.Properties != nil && req.Properties.Name != nil {
		updateData["name"] = *req.Properties.Name
	}

	// update geometry
	if req.Geometry != nil {
		updateData["location"] = req.Geometry
	}

	if req.Properties != nil && req.Properties.Category != nil {
		updateData["category"] = *req.Properties.Category
	}

	if len(updateData) == 0 {
		return c.JSON(400, map[string]string{
			"error": "no fields to update",
		})
	}

	if err := pc.Service.Update(c.Request().Context(), id, updateData); err != nil {
		return c.JSON(500, map[string]string{
			"error": "update failed",
		})
	}

	return c.JSON(200, map[string]string{
		"message": "patched successfully",
	})
}

type PolygonGeometry struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

type WithinRequest struct {
	Polygon PolygonGeometry `json:"polygon"`
}

func (pc *PlaceController) FindWithin(c echo.Context) error {
	var req WithinRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request",
		})
	}

	places, err := pc.Service.FindWithin(
		c.Request().Context(),
		req.Polygon, // ส่งตรงไปเลย
	)

	if err != nil {
		return c.JSON(500, map[string]string{
			"error": "failed to query",
		})
	}

	features := make([]entity.Feature, 0)
	for _, p := range places {
		features = append(features, toFeature(p))
	}

	return c.JSON(200, entity.FeatureCollection{
		Type:     "FeatureCollection",
		Features: features,
	})
}