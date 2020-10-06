package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"wb2-master/api/models"
)

var roles = []models.Role{
	models.Role{
		Name: "Admin",
	},
	models.Role{
		Name: "Mitra",
	},
	models.Role{
		Name: "Agent",
	},
	models.Role{
		Name: "Sub Agent",
	},
}
var users = []models.User{
	models.User{
		Username: "admin",
		Email:    "admin@me.com",
		Password: "Pandi123#",
	},
}
func Load(db *gorm.DB) {
	var err error
	//role
	for i, _ := range roles {
		err = db.Debug().Model(&models.Role{}).Create(&roles[i]).Error
		if err != nil {
			log.Fatalf("cannot seed roles table: %v", err)
		}
	}

	//user
	for i, _ := range users {
		users[i].RoleID = 1
	
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

	}
}