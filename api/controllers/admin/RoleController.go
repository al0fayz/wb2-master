package admin

import (
	"github.com/gofiber/fiber/v2"

	"wb2-master/api/databases"
	"wb2-master/api/models"
)

//get all Roles
func GetAllRole(c *fiber.Ctx) error {
	db := databases.DB
	var roles []models.Role
	db.Find(&roles)	//get all data
	return c.JSON(fiber.Map{"status": "success", "message": "All products", "data": roles})
}

//get role by ID
func GetRole(c *fiber.Ctx) error {
	id := c.Params("id")
	db := databases.DB
	var role models.Role
	db.Find(&role, id)
	if role.Name == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No role found with ID", "data": nil})

	}
	return c.JSON(fiber.Map{"status": "success", "message": "Role found", "data": role})
}
//save role
func SaveRole(c *fiber.Ctx) error {
	var err error
	db := databases.DB
	role := new(models.Role)

	if err := c.BodyParser(role); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create role", "data": err})
	}
	role.Prepare()
	err = role.Validate()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Validate error", "data": err})
	}
	db.Create(&role)
	return c.JSON(fiber.Map{"status": "success", "message": "Created product", "data": role})
}