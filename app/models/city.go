package models

import (
	"time"

	"github.com/emvi/hide"
)

// City model
type City struct {
	ID hide.ID `gorm:"primary_key" json:"id"`

	Name string `json:"name" valid:"required"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
