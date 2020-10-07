package router

import (
	"wb2-master/api/controllers"
	"wb2-master/api/controllers/admin"

	// "wb2-master/api/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", controllers.Home)

	//role
	role := api.Group("/role")
	role.Get("/", admin.GetAllRole)
	role.Post("/", admin.SaveRole)
	role.Get("/:id", admin.GetRole)
	role.Delete("/:id", admin.DeleteRole)
}