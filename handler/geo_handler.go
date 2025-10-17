package handler

import (
	"context"
	"fmt"
	"inibackend/model"
	"inibackend/repository"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func timeoutCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

// POST /geo
func CreateGeo(c *fiber.Ctx) error {
	var fc model.FeatureCollection
	if err := c.BodyParser(&fc); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid body", "detail": err.Error()})
	}
	if fc.Type == "" {
		fc.Type = "FeatureCollection"
	}
	ctx, cancel := timeoutCtx()
	defer cancel()
	id, err := repository.CreateFeatureCollection(ctx, fc)
	if err != nil {
		fmt.Printf("CreateGeo: failed to insert: %v\n", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	fmt.Printf("CreateGeo: inserted with id: %v\n", id)
	return c.Status(http.StatusCreated).JSON(fiber.Map{"insertedId": id})
}

// POST /geo/bulk accepts a FeatureCollection and inserts all features
func CreateGeoBulk(c *fiber.Ctx) error {
	var fc model.FeatureCollection
	if err := c.BodyParser(&fc); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid body", "detail": err.Error()})
	}
	if fc.Type == "" {
		fc.Type = "FeatureCollection"
	}
	ctx, cancel := timeoutCtx()
	defer cancel()
	id, err := repository.CreateFeatureCollection(ctx, fc)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"insertedId": id})
}

// GET /geo
func ListGeo(c *fiber.Ctx) error {
	ctx, cancel := timeoutCtx()
	defer cancel()
	items, err := repository.ListFeatureCollections(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(items)
}

// GET /geo/:id
func GetGeo(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx, cancel := timeoutCtx()
	defer cancel()
	item, err := repository.GetFeatureCollection(ctx, id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(item)
}

// PUT /geo/:id
func UpdateGeo(c *fiber.Ctx) error {
	id := c.Params("id")
	var fc model.FeatureCollection
	if err := c.BodyParser(&fc); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid body", "detail": err.Error()})
	}
	if fc.Type == "" {
		fc.Type = "FeatureCollection"
	}
	ctx, cancel := timeoutCtx()
	defer cancel()
	n, err := repository.UpdateFeatureCollection(ctx, id, fc)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"modified": n})
}

// DELETE /geo/:id
func DeleteGeo(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx, cancel := timeoutCtx()
	defer cancel()
	n, err := repository.DeleteFeatureCollection(ctx, id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if n == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(fiber.Map{"deleted": n})
}

// GET /api/test-db
func TestDB(c *fiber.Ctx) error {
	ctx, cancel := timeoutCtx()
	defer cancel()
	col, err := repository.GeoCol()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "error": err.Error()})
	}
	err = col.Database().Client().Ping(ctx, nil)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "error": err.Error()})
	}
	return c.JSON(fiber.Map{"status": "connected"})
}
