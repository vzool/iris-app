package models

import (
	"database/sql"
	"time"

	"github.com/emvi/hide"
)

// Task model
type Task struct {
	ID hide.ID `gorm:"primary_key" json:"id"`

	Name      string        `json:"name" valid:"required"`
	CityID    sql.NullInt64 `json:"city_id" valid:"required,int" gorm:"DEFAULT:null;foreignkey:CityID"`
	Completed bool          `json:"completed" gorm:"DEFAULT:false"`
	City      City          `json:"city"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
