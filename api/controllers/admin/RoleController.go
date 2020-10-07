package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/biezhi/gorm-paginator/pagination"
	"strconv"
	"wb2-master/api/databases"
	"wb2-master/api/models"
)

//get all Roles
func GetAllRole(c *fiber.Ctx) error {
	db := databases.DB
	var page = 1
	var limit = 10
	var roles []models.Role

	if val, err := strconv.Atoi(c.Query("page")); err == nil {
		page = val
	}

	if val, err := strconv.Atoi(c.Query("limit")); err == nil {
		limit = val
	}
	//search
	qSearch := c.Query("q")
	if len(qSearch) > 0 {
		db = db.Where("name LIKE ?", "%"+qSearch+"%")
	}
	//pagination

	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"created_at desc"},
		ShowSQL: true,
	}, &roles)

	return c.JSON(fiber.Map{
		"success": true,
		"status": "success", 
		"messages": "All products", 
		"data": paginator,
	})
}

//get role by ID
func GetRole(c *fiber.Ctx) error {
	id := c.Params("id")
	db := databases.DB
	var role models.Role
	db.Find(&role, id)
	if role.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"status": "error", 
			"messages": "No role found with ID", 
			"data": nil,
		})

	}
	return c.JSON(fiber.Map{
		"success": true,
		"status": "success", 
		"messages": "Role found", 
		"data": role,
	})
}
//save role
func SaveRole(c *fiber.Ctx) error {
	var err error
	db := databases.DB
	role := new(models.Role)

	if err = c.BodyParser(role); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"success": false,
			"status": "error", 
			"messages": err.Error(), 
			"data": nil,
		})
	}
	role.Prepare()
	//validate name 
	err = role.Validate()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"status": "error", 
			"messages": err.Error(), 
			"data": nil,
		})
	}
	db.Create(&role)
	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "Created product", 
		"data": role,
	})
}
