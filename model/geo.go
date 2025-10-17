package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Geometry represents a GeoJSON geometry object
type Geometry struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}

// Feature represents a GeoJSON Feature
type Feature struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type       string             `json:"type" bson:"type"`
	Properties map[string]any     `json:"properties" bson:"properties"`
	Geometry   Geometry           `json:"geometry" bson:"geometry"`
}

// FeatureCollection represents a collection of GeoJSON Features
type FeatureCollection struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type     string             `json:"type" bson:"type"`
	Features []Feature          `json:"features" bson:"features"`
}
