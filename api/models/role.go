package models

import (
	"html"
	"strings"
	"time"
	"github.com/go-playground/validator/v10"
)

type Role struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name	  string    `gorm:"size:255;not null;unique" json:"name" validate:"required"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (r *Role) Prepare() {
	r.ID 		= 0
	r.Name 		= html.EscapeString(strings.TrimSpace(r.Name))
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}
func (r *Role) Validate() error {
	var validate *validator.Validate
	validate = validator.New()
	err := validate.Struct(r)
	if err != nil {
		return err
	}
	return nil
}