package seed

import (
	"wb2-master/api/utils/password"
	"log"

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
var users = []models.User{
	models.User{
		Username: "admin",
		Email:    "admin@me.com",
		Password: password.Generate("Pandi123#"),
		IsActive: true,
		IsAdmin: true,
		RoleID: 1,
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
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

	}
}