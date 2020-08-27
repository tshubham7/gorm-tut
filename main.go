package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tshubham7/gorm-articles/database"
	"github.com/tshubham7/gorm-articles/handler"
	"github.com/tshubham7/gorm-articles/models"
	"github.com/tshubham7/gorm-articles/repository"
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

	r := gin.Default()

	u := repository.NewUserService(db)
	auth(r, u)

	r.Run(":8080")
}

// auth service routes
func auth(r *gin.Engine, u repository.UserService) {
	auth := handler.NewAuthHandler(u)

	r.POST("/auth/register", auth.Register())

}
