package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

// Task model
type Task struct {
	gorm.Model
	Name      string        `json:"name" valid:"required"`
	CityID    sql.NullInt64 `json:"city_id" valid:"required,int" gorm:"DEFAULT:null;foreignkey:CityID"`
	Completed bool          `json:"completed" gorm:"DEFAULT:false"`
	City      City          `json:"city"`
}
