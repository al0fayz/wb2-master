package migration

import (
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	"wb2-master/api/models"
)

func Migrate(db *gorm.DB) {
	//delete table
	err := db.DropTableIfExists(
		&models.UserDetail{}, 
		&models.Country{}, 
		&models.User{}, 
		&models.Role{}).Error

	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	fmt.Println("Success Drop Tables")
	//migrate table
	err = db.AutoMigrate(
		&models.Role{}, 
		&models.User{}, 
		&models.Country{}, 
		&models.UserDetail{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
	fmt.Println("Success Migrate Tables")
	//id foreign key role id in user table
	err = db.Model(&models.User{}).AddForeignKey("role_id", "roles(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	//id foreign key user id in userdetail table
	err = db.Model(&models.UserDetail{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	//id foreign key country id in userdetail table
	err = db.Model(&models.UserDetail{}).AddForeignKey("country_id", "country(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	fmt.Println("Success attaching foreign key")
}