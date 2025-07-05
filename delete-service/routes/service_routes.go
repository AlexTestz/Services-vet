package routes

import (
	"github.com/AlexTestz/delete-service/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/services")
	api.Delete("/:id", controllers.DeleteService)
	api.Delete("/name/:name", controllers.DeleteServiceByName) //
}
