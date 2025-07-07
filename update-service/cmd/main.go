package main

import (
	"github.com/AlexTestz/update-service/config"
	"github.com/AlexTestz/update-service/database"
	"github.com/AlexTestz/update-service/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
//	config.LoadEnv()
	database.ConnectMongo()

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Next()
	})

	routes.SetupRoutes(app)

	app.Listen(":" + config.GetEnv("PORT"))
	app.Use(cors.New(cors.Config{
  AllowOrigins: "*",  // Permite solicitudes desde cualquier origen
  AllowMethods: "GET, POST, PUT, DELETE",  // Permite PUT

}))

}

