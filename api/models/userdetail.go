package models

import (
	"time"
)

type UserDetail struct {
	ID				uint32    `gorm:"primary_key;auto_increment" json:"id"`
	FirstName		string    `gorm:"size:255;not null;" json:"first_name"`
	LastName     	string    `gorm:"size:255;not null;" json:"last_name"`
	Organization  	string    `gorm:"size:100;not null;" json:"organization"`
	State		  	string    `gorm:"size:255;not null;" json:"state"`
	City  			string    `gorm:"size:255;not null;" json:"city"`
	Street  		string    `gorm:"size:255;not null;" json:"street"`
	PostalCode  	string    `gorm:"size:100;not null;" json:"postal_code"`
	Phone		  	string    `gorm:"size:100;not null;" json:"phone"`
	Fax			  	string    `gorm:"size:100;" json:"fax"`
	Images			string    `gorm:"size:100;" json:"images"`
	UserID    		uint32    `gorm:"not null" json:"user_id"`
	User      		User      `json:"user"`
	CountryID    	uint32    `gorm:"not null" json:"countri_id"`
	Country      	Country   `json:"country"`
	CreatedAt 		time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 		time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt 		*time.Time `json:"deleted_at"`
}
