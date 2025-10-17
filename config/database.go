package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database and collection configuration
// Adjusted to match user's request: DB "geo" and collection "geojson"
var DBName = func() string {
	if v := os.Getenv("DB_NAME"); v != "" {
		return v
	}
	return "geo"
}()
var MahasiswaCollection = "geojson"

// Connection string: prefer env MONGOSTRING, MONGODB_URL, or DATABASE_URL if present, otherwise fall back to provided Atlas URI
var MongoString string = func() string {
	if v := os.Getenv("MONGOSTRING"); v != "" {
		return v
	}
	if v := os.Getenv("MONGODB_URL"); v != "" {
		return v
	}
	if v := os.Getenv("DATABASE_URL"); v != "" {
		return v
	}
	return "mongodb+srv://xxalvaro158:Faridalvaro158@webservice.nb18pin.mongodb.net/?retryWrites=true&w=majority&appName=WebService"
}()
var UserCollection = "user"

func MongoConnect(dbname string) (*mongo.Database, error) {
	fmt.Printf("MongoConnect: using connection string: %s\n", MongoString)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}
	return client.Database(dbname), nil
}
