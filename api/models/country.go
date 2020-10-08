package models

import (
	"time"
)

type Country struct {
	ID        	uint32    	`gorm:"primary_key;auto_increment" json:"id"`
	Code 	  	string 		`gorm:"size:100;not null;" json:"code"`
	Name	  	string    	`gorm:"size:255;not null;unique" json:"name"`
	PhoneCode	uint32 		`gorm:"not null;" json:"phonecode"`
	CreatedAt 	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
type Tabler interface {
	TableName() string
  }
  
  // TableName overrides the table name used by User to `profiles`
  func (Country) TableName() string {
	return "country"
  }