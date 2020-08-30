package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tshubham7/gorm-articles/repository"
	"github.com/tshubham7/gorm-articles/services"
)

type article struct {
	a repository.ArticleRepo
}

// ArticleHandler ...
type ArticleHandler interface {
	// create article
	Create() gin.HandlerFunc

	// list articles
	List() gin.HandlerFunc
}

// NewArticleHandler ...
func NewArticleHandler(a repository.ArticleRepo) ArticleHandler {
	return &article{a}
}

// Create ...
func (ar article) Create() gin.HandlerFunc {
	sr := services.NewArticleService(ar.a)

	return func(c *gin.Context) {
		// c.Request
		var params services.ArticleCreateRequest
		err := c.Bind(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "missing or invalid params",
				"error":   err.Error(),
			})
			return
		}

		article, err := sr.Create(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to create article",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, article)
	}
}

// List ...
func (ar article) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		// userID := mid.UserID(c)
		// fmt.Println(userID)
		fmt.Println("list api called")
	}
}
