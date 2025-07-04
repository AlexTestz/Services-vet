package routes

import (
	"github.com/AlexTestz/update-service/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/services")
	api.Put("/:id", controllers.UpdateService)
	    api.Put("/name/:name", controllers.UpdateServiceByName) // Nueva ruta para actualizar por nombre

}
