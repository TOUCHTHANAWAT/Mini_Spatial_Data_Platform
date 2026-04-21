package entity

type Feature struct {
	Type       string      `json:"type"` // "Feature"
	ID         string      `json:"id"`
	Geometry   Geometry    `json:"geometry"`
	Properties interface{} `json:"properties"`
}

type FeatureCollection struct {
	Type     string    `json:"type"` // "FeatureCollection"
	Features []Feature `json:"features"`
}