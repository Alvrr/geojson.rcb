package router

import (
	"inibackend/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	geo := api.Group("/geo")

	geo.Post("/", handler.CreateGeo)
	geo.Get("/", handler.ListGeo)
	geo.Get("/:id", handler.GetGeo)
	geo.Put("/:id", handler.UpdateGeo)
	geo.Delete("/:id", handler.DeleteGeo)

	// Test DB connection
	api.Get("/test-db", handler.TestDB)
}
