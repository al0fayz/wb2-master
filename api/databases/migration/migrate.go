package migration

import (
	"log"
	"github.com/jinzhu/gorm"
	"wb2-master/api/models"
)

func Migrate(db *gorm.DB) {
	//delete table
	err := db.Debug().DropTableIfExists(&models.User{}, &models.Role{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	
	//migrate table
	err = db.Debug().AutoMigrate(&models.Role{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	//id foreign key role id in user table
	err = db.Debug().Model(&models.User{}).AddForeignKey("role_id", "roles(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

}