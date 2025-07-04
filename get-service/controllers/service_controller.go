package controllers

import (
	"context"
	"time"

	"github.com/AlexTestz/get-service/database"
	"github.com/AlexTestz/get-service/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetServices(c *fiber.Ctx) error {
    collection := database.Client.Database("pet_care_services_db").Collection("services")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // ✅ Leer query param
    name := c.Query("name")

    // 📄 Condición de búsqueda
    filter := bson.M{}
    if name != "" {
        filter = bson.M{"name": bson.M{"$regex": name, "$options": "i"}} // búsqueda insensible a mayúsculas
    }

    cursor, err := collection.Find(ctx, filter)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching services"})
    }

    var services []models.Service
    if err := cursor.All(ctx, &services); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error parsing results"})
    }

    return c.JSON(services)
}
type Service struct {
    ID          string `json:"id" bson:"_id,omitempty"`
    Name        string `json:"name" bson:"name"`
    Description string `json:"description" bson:"description"`
    Price       int    `json:"price" bson:"price"`
    Duration    int    `json:"duration" bson:"duration"`
}

func GetServiceByName(c *fiber.Ctx) error {
    name := c.Params("name")
    if name == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name is required"})
    }

    collection := database.Client.Database("pet_care_services_db").Collection("services")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var service Service
    filter := bson.M{"name": bson.M{"$regex": name, "$options": "i"}} // <-- búsqueda insensible a mayúsculas/minúsculas
    err := collection.FindOne(ctx, filter).Decode(&service)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Service not found"})
    }

    return c.JSON(service)
}