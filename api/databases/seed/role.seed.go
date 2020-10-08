package seed

import (
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	"wb2-master/api/models"
)

var roles = []models.Role{
	models.Role{
		ID: 1,
		Name: "Admin",
	},
	models.Role{
		ID: 2,
		Name: "Mitra",
	},
	models.Role{
		ID: 3,
		Name: "Agent",
	},
	models.Role{
		ID: 4,
		Name: "Sub Agent",
	},
}
func RoleSeed(db *gorm.DB) {
	var err error
	//role
	for i, _ := range roles {
		err = db.Model(&models.Role{}).Create(&roles[i]).Error
		if err != nil {
			log.Fatalf("cannot seed roles table: %v", err)
		}
	}
	fmt.Println("Success seed role table")
}