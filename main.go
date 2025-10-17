package main

import (
	"context"
	"encoding/json"
	"fmt"
	"inibackend/config"
	"inibackend/model"
	"inibackend/repository"
	r "inibackend/router"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load environment variables: ", err)
	}
}

func loadGeoJSONData() {
	// Read the kampung.geojson file
	data, err := ioutil.ReadFile("kampung.geojson")
	if err != nil {
		fmt.Printf("Failed to read kampung.geojson: %v\n", err)
		return
	}

	// Parse the JSON data into FeatureCollection
	var fc model.FeatureCollection
	err = json.Unmarshal(data, &fc)
	if err != nil {
		fmt.Printf("Failed to parse GeoJSON: %v\n", err)
		return
	}

	// Insert into database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := repository.CreateFeatureCollection(ctx, fc)
	if err != nil {
		fmt.Printf("Failed to insert GeoJSON data into DB: %v\n", err)
		return
	}

	fmt.Printf("Successfully loaded GeoJSON data from kampung.geojson into DB with ID: %v\n", id)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic recovered: %v\n", r)
			log.Fatalf("Application panicked: %v", r)
		}
	}()
	fmt.Println("Starting application...")
	// Test DB connection on startup (optional, log but continue if fail)
	_, err := config.MongoConnect("geo")
	if err != nil {
		fmt.Printf("Warning: Failed to connect to DB on startup: %v\n", err)
		fmt.Println("Continuing without DB connection...")
	} else {
		fmt.Println("DB connection successful on startup")
		// Load GeoJSON data from kampung.geojson
		loadGeoJSONData()
	}
	// Initialize the Fiber app
	app := fiber.New()
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: false,
		AllowMethods:     "GET,POST,PUT,DELETE",
	}))

	r.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "Endpoint not found",
		})

	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8088" // Default port if not set in .env file
	}
	fmt.Printf("Server is running on port %s\n", port)

	// Start server in a goroutine
	go func() {
		if err := app.Listen("0.0.0.0:" + port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Gracefully shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
