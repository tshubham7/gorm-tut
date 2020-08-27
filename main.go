package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tshubham7/gorm-articles/database"
	"github.com/tshubham7/gorm-articles/models"
)

func main() {
	db, err := database.ConnectToDB("postgres", "1234", "gorm_article")
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Article{})
	db.AutoMigrate(&models.Comment{})
	fmt.Println("migration success")
}
