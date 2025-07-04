package main

import (
	"github.com/AlexTestz/create-service/config"
	"github.com/AlexTestz/create-service/database"
	"github.com/AlexTestz/create-service/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
    config.LoadEnv()
    database.ConnectMongo()

    app := fiber.New()

    app.Use(func(c *fiber.Ctx) error {
        c.Set("Content-Type", "application/json")
        return c.Next()
    })

    routes.SetupRoutes(app)

    app.Listen(":" + config.GetEnv("PORT"))
}
