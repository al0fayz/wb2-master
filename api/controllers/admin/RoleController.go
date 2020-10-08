package admin

import (
	"wb2-master/api/databases"
	"wb2-master/api/models"
	"wb2-master/api/utils"
	
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
)
type RoleInput struct {
	Name string `json:"name" validate:"required,alpha"`
}
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
		"messages": "All role", 
		"data": paginator,
	})
}

//get role by ID
func GetRole(c *fiber.Ctx) error {
	id := c.Params("id")
	db := databases.DB
	var role models.Role
	err := db.Find(&role, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return fiber.NewError(404, "Role not found!")
		}

	}
	return c.JSON(fiber.Map{
		"success": true,
		"status": "success", 
		"messages": "Data Found", 
		"data": role,
	})
}
//save role
func SaveRole(c *fiber.Ctx) error {
	db := databases.DB
	roleinput := new(RoleInput)
	role := models.Role{}

	if err := utils.ParseBodyAndValidate(c, roleinput); err != nil {
		return err
	}
	role.Name = roleinput.Name
	role.Prepare()
	
	if err := db.Create(&role).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}
	return c.JSON(fiber.Map{
		"success": true,
		"status": "success", 
		"messages": "Successfully save data", 
		"data": role,
	})
}
// Update Role
func UpdateRole(c *fiber.Ctx) error {
	db := databases.DB
	roleinput := new(RoleInput)
	role := models.Role{}

	if err := utils.ParseBodyAndValidate(c, roleinput); err != nil {
		return err
	}
	
	id := c.Params("id")

	err := db.First(&role, id).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return fiber.NewError(404, "Role not found!")
		}
	}
	role.Name = roleinput.Name
	
	db.Save(&role)

	return c.JSON(fiber.Map{
		"success": true,
		"status": "success", 
		"messages": "Successfully update data", 
		"data": role,
	})
}

// Delete Role
func DeleteRole(c *fiber.Ctx) error {
	id := c.Params("id")
	db := databases.DB

	var role models.Role
	err := db.First(&role, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return fiber.NewError(404, "Role not found!")
		}
	}
	db.Delete(&role)
	return c.JSON(fiber.Map{
		"success": true,
		"status": "success", 
		"message": "Data successfully deleted", 
		"data": nil,
	})
}