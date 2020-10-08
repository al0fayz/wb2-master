package routes

import (
	"wb2-master/api/controllers/subagent"
	"wb2-master/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SubAgentRoutes(app fiber.Router) {
	//must login
	v2 := app.Group("/api/v2").Use(middlewares.Auth)

	//agent
	subagentRole := v2.Group("/subagent").Use(middlewares.IsSubAgent)
	
	//dashboard 
	subagentRole.Get("/dashboard", subagent.Dashboard)
	
}