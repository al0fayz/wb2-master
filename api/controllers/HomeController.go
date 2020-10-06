package controllers

import "github.com/gofiber/fiber/v2"

// Hello hanlde api status
func Home(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}
