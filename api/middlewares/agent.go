package middlewares

import (
	"wb2-master/api/utils"
	"wb2-master/api/models"
	"wb2-master/api/databases"
	
	"github.com/gofiber/fiber/v2"
)

// Auth is the authentication middleware
func IsAgent(c *fiber.Ctx) error {
	id := utils.GetUser(c)
	//search user
	var u models.User
	err := databases.DB.Find(&u, *id).Error
	if err != nil {
		return fiber.ErrUnauthorized
	}
	if u.RoleID != uint32(3) {
		return fiber.ErrUnauthorized
	}
	return c.Next()
}
