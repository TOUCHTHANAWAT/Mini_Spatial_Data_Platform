package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Geometry struct {
	Type        string     `json:"type" bson:"type"` // Point
	Coordinates [2]float64 `json:"coordinates" bson:"coordinates"`
}

type Place struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Location Geometry           `json:"location" bson:"location"`
}