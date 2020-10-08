package routes

import (
	"wb2-master/api/controllers"
	"wb2-master/api/controllers/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Home
	api := app.Group("/api")
	api.Get("/", controllers.Home)
	//login
	api.Post("/login", auth.Login)
}