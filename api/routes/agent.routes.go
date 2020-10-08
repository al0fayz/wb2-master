package routes

import (
	"wb2-master/api/controllers/agent"
	"wb2-master/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AgentRoutes(app fiber.Router) {
	//must login
	v2 := app.Group("/api/v2").Use(middlewares.Auth)

	//agent
	agentRole := v2.Group("/agent").Use(middlewares.IsAgent)
	
	//dashboard 
	agentRole.Get("/dashboard", agent.Dashboard)
	
}