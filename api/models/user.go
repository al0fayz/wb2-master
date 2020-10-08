package models

import (
	"time"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	IsActive  bool		`gorm:"default:false" json:"isActive"`
	IsAdmin   bool		`gorm:"default:false" json:"isAdmin"`
	RoleID    uint32    `gorm:"not null" json:"role_id"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
