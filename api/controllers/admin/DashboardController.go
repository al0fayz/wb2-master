package admin

import "github.com/gofiber/fiber/v2"

// Hello hanlde api status
func Dashboard(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "Hello Admin i'm ok!", 
		"data": nil,
	})
}
