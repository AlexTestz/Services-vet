package controllers

import (
	"context"
	"time"

	"github.com/AlexTestz/delete-service/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteService(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	collection := database.Client.Database("pet_care_services_db").Collection("services")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete service"})
	}

	if res.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Service not found"})
	}

	return c.JSON(fiber.Map{"message": "✅ Service deleted successfully"})
}

func DeleteServiceByName(c *fiber.Ctx) error {
    name := c.Params("name")
    if name == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name is required"})
    }

    collection := database.Client.Database("pet_care_services_db").Collection("services")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    res, err := collection.DeleteOne(ctx, bson.M{"name": name})
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete service"})
    }

    if res.DeletedCount == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Service not found"})
    }

    return c.JSON(fiber.Map{"message": "✅ Service deleted successfully"})
}