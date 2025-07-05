package main

import (
	"github.com/AlexTestz/get-service/config"
	"github.com/AlexTestz/get-service/database"
	"github.com/AlexTestz/get-service/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
   // config.LoadEnv()
    database.ConnectMongo()

    app := fiber.New()

    app.Use(func(c *fiber.Ctx) error {
        c.Set("Content-Type", "application/json")
        return c.Next()
    })

    routes.SetupRoutes(app)

    app.Listen(":" + config.GetEnv("PORT"))
}
