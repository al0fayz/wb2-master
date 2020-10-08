package router

import (
	"wb2-master/api/controllers"
	"wb2-master/api/controllers/admin"
	"wb2-master/api/controllers/auth"

	"wb2-master/api/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Home
	api := app.Group("/api", logger.New())
	v2 := api.Group("/v2")
	v2.Get("/", controllers.Home)

	//auth
	v2.Post("/login", auth.Login)

	//admin
	adminRole := v2.Group("/admin", middlewares.Protected())
	
	//role
	role := adminRole.Group("/role")
	role.Get("/", admin.GetAllRole)
	role.Post("/", admin.SaveRole)
	role.Get("/:id", admin.GetRole)
	role.Put("/:id", admin.UpdateRole)
	role.Delete("/:id", admin.DeleteRole)
}