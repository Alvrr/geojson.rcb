package repository

import (
	"context"
	"fmt"
	"inibackend/config"
	"inibackend/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func geoCol() (*mongo.Collection, error) {
	db, err := config.MongoConnect(config.DBName)
	if err != nil {
		return nil, err
	}
	return db.Collection(config.MahasiswaCollection), nil
}

func GeoCol() (*mongo.Collection, error) {
	return geoCol()
}

func CreateFeatureCollection(ctx context.Context, fc model.FeatureCollection) (interface{}, error) {
	col, err := geoCol()
	if err != nil {
		return nil, err
	}
	fmt.Printf("CreateFeatureCollection: attempting to insert into DB %s, collection %s\n", config.DBName, config.MahasiswaCollection)
	res, err := col.InsertOne(ctx, fc)
	if err != nil {
		fmt.Printf("CreateFeatureCollection: insert error: %v\n", err)
		return nil, err
	}
	fmt.Printf("CreateFeatureCollection: insert successful, id: %v\n", res.InsertedID)
	return res.InsertedID, nil
}

func ListFeatureCollections(ctx context.Context) ([]model.FeatureCollection, error) {
	col, err := geoCol()
	if err != nil {
		return nil, err
	}
	// Find all FeatureCollection documents
	cur, err := col.Find(ctx, bson.M{"type": "FeatureCollection"})
	if err != nil {
		return nil, err
	}
	var out []model.FeatureCollection
	if err := cur.All(ctx, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func GetFeatureCollection(ctx context.Context, id string) (model.FeatureCollection, error) {
	col, err := geoCol()
	if err != nil {
		return model.FeatureCollection{}, err
	}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.FeatureCollection{}, err
	}
	var fc model.FeatureCollection
	err = col.FindOne(ctx, bson.M{"_id": oid}).Decode(&fc)
	return fc, err
}

func UpdateFeatureCollection(ctx context.Context, id string, fc model.FeatureCollection) (int64, error) {
	col, err := geoCol()
	if err != nil {
		return 0, err
	}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	update := bson.M{"$set": bson.M{
		"type":     fc.Type,
		"features": fc.Features,
	}}
	res, err := col.UpdateByID(ctx, oid, update)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}

func DeleteFeatureCollection(ctx context.Context, id string) (int64, error) {
	col, err := geoCol()
	if err != nil {
		return 0, err
	}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	res, err := col.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}
