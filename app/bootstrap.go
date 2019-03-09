package app

import (
	"github.com/emvi/hide"
	"github.com/jinzhu/gorm"

	// SQLite dialect for Gorm
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// SecretKey to secure the App
func SecretKey() string {
	return "941C1A0A-AF24-42F5-96D8-A7092754" // Must be 32 bytes
}

func init() {
	// setup hideID and here we need some secret
	hide.UseHash(hide.NewHashID(SecretKey(), 9))
}

// DB Connection to database
func DB() *gorm.DB {

	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("failed to connect database")
	}

	db.LogMode(true)
	db.SingularTable(true)

	return db
}
