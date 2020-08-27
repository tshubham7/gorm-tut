package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// ConnectToDB ...
func ConnectToDB(dbUser, dbPassword, dbName string) (*gorm.DB, error) {
	var connString = fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName,
	)
	db, err := gorm.Open("postgres", connString)
	// failed to connect
	if err != nil {
		return db, err
	}

	err = db.DB().Ping()
	return db, err
}
