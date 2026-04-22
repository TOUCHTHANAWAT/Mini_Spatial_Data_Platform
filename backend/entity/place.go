package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Geometry struct {
	Type        string     `json:"type" bson:"type"` // Point
	Coordinates [2]float64 `json:"coordinates" bson:"coordinates"`
}

type Place struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Category  string             `json:"category" bson:"category"`
	Location Geometry           `json:"location" bson:"location"`
	DeletedAt *time.Time         `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
}