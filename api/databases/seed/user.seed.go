package seed

import (
	"wb2-master/api/utils/password"
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	"wb2-master/api/models"
)

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
func UserSeed(db *gorm.DB) {
	var err error
	//user
	for i, _ := range users {
		err = db.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
	fmt.Println("Success seed users table")
}