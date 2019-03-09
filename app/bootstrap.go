package app

import (
	"github.com/emvi/hide"
	"github.com/jinzhu/gorm"

	// SQLite dialect for Gorm
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func init() {
	// setup hideID and here we need some secret
	hide.UseHash(hide.NewHashID("iris-app", 9))
}

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
