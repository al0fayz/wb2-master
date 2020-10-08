package routes

import (
	"wb2-master/api/controllers/admin"
	"wb2-master/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(app fiber.Router) {
	//must login
	v2 := app.Group("/api/v2").Use(middlewares.Auth)

	//admin
	adminRole := v2.Group("/admin")
	
	//role
	role := adminRole.Group("/role")
	role.Get("/", admin.GetAllRole)
	role.Post("/", admin.SaveRole)
	role.Get("/:id", admin.GetRole)
	role.Put("/:id", admin.UpdateRole)
	role.Delete("/:id", admin.DeleteRole)
}