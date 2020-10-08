package agent

import "github.com/gofiber/fiber/v2"

// Hello hanlde api status
func Dashboard(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "Hello Agent i'm ok!", 
		"data": nil,
	})
}
