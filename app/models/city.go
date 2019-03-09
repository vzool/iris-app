package models

import "github.com/jinzhu/gorm"

// City model
type City struct {
	gorm.Model
	Name string `json:"name" valid:"required"`
}
