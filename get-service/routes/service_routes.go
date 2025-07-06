package routes

import (
	"github.com/AlexTestz/get-service/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api/services")
    api.Get("/", controllers.GetServices)
	api.Get("/name/:name", controllers.GetServiceByName)
	api.Get(":id", controllers.GetServiceByID)
}
