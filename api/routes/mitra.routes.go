package routes

import (
	"wb2-master/api/controllers/mitra"
	"wb2-master/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func MitraRoutes(app fiber.Router) {
	//must login
	v2 := app.Group("/api/v2").Use(middlewares.Auth)

	//mitra
	mitraRole := v2.Group("/mitra").Use(middlewares.IsMitra)
	
	//dashboard 
	mitraRole.Get("/dashboard", mitra.Dashboard)
	
}