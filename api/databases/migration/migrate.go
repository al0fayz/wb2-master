package migration

import (
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	"wb2-master/api/models"
)

//delete all table
func DropTable(db *gorm.DB) {
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
}

func Migrate(db *gorm.DB) {
	//check table Role
	role := db.HasTable(&models.Role{})
	if !role {
		err := db.CreateTable(&models.Role{}).Error
		if err != nil {
			log.Fatalf("cannot migrate table: %v", err)
		}
		fmt.Println("Success Migrate Tables Role")
	}
	//check table User
	user := db.HasTable(&models.User{})
	if !user {
		err := db.CreateTable(&models.User{}).Error
		if err != nil {
			log.Fatalf("cannot migrate table: %v", err)
		}
		
		//id foreign key role id in user table
		err = db.Model(&models.User{}).AddForeignKey("role_id", "roles(id)", "cascade", "cascade").Error
		if err != nil {
			log.Fatalf("attaching foreign key error: %v", err)
		}
		fmt.Println("Success Migrate Tables User")
	}
	//check table Country
	country := db.HasTable(&models.Country{})
	if !country {
		err := db.CreateTable(&models.Country{}).Error
		if err != nil {
			log.Fatalf("cannot migrate table: %v", err)
		}
		fmt.Println("Success Migrate Tables Country")
	}
	//check table Detail User
	ud := db.HasTable(&models.UserDetail{})
	if !ud {
		err := db.CreateTable(&models.UserDetail{}).Error
		if err != nil {
			log.Fatalf("cannot migrate table: %v", err)
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
		fmt.Println("Success Migrate Tables User Detail")
	}
	return
}