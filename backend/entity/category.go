package entity

type Category struct {
	Value string `json:"value" bson:"value"`
	Label string `json:"label" bson:"label"`
	Color string `json:"color" bson:"color"`
}