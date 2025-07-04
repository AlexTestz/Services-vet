package controllers

import (
	"context"
	"time"

	"github.com/AlexTestz/create-service/database"
	"github.com/AlexTestz/create-service/models"
	"github.com/gofiber/fiber/v2"
)

func CreateService(c *fiber.Ctx) error {
    collection := database.Client.Database("pet_care_services_db").Collection("services")

    var service models.Service
    if err := c.BodyParser(&service); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    // 🔍 Validaciones manuales
    if service.Name == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name is required"})
    }
    if service.Price <= 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Price must be greater than 0"})
    }
    if service.Duration <= 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Duration must be greater than 0"})
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    res, err := collection.InsertOne(ctx, service)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert"})
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Service created ✅",
        "id":      res.InsertedID,
    })
}
