package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tshubham7/gorm-articles/database"
	"github.com/tshubham7/gorm-articles/handler"
	mid "github.com/tshubham7/gorm-articles/middleware"
	"github.com/tshubham7/gorm-articles/models"
	"github.com/tshubham7/gorm-articles/repository"
)

func main() {
	// use environment variables
	db, err := database.ConnectToDB("postgres", "1234", "gorm_article")

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Article{})
	db.AutoMigrate(&models.Comment{})

	r := gin.Default()
	u := repository.NewUserRepo(db)
	ar := repository.NewArticleRepo(db)
	auth(r, u)
	article(r, ar)
	r.Run(":8080")
}

// auth service routes
func auth(r *gin.Engine, u repository.UserRepo) {
	h := handler.NewAuthHandler(u)
	route := r.Group("api/auth/")
	{
		route.POST("register", h.Register())
		route.POST("login", h.Login())
	}
}

// article service routes
func article(r *gin.Engine, a repository.ArticleRepo) {
	h := handler.NewArticleHandler(a)
	route := r.Group("api/article")

	route.Use(mid.Authenticator()) // jwt auth header
	{
		route.POST("", h.Create())
		route.GET("", h.List())
	}
}
