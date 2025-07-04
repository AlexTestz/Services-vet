package routes

import (
	"github.com/AlexTestz/create-service/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api/services")
    api.Post("/", controllers.CreateService)
}
