package main

import (
	"github.com/AlexTestz/delete-service/config"
	"github.com/AlexTestz/delete-service/database"
	"github.com/AlexTestz/delete-service/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
//	config.LoadEnv()
	database.ConnectMongo()

	app := fiber.New()
	routes.SetupRoutes(app)

	app.Listen(":" + config.GetEnv("PORT"))
}
