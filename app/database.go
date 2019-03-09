package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB Connection
func DB() *gorm.DB {

	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("failed to connect database")
	}

	db.LogMode(true)
	db.SingularTable(true)

	return db
}
