package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mid "github.com/tshubham7/gorm-articles/middleware"
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
		var params services.ArticleCreateRequest
		var err error

		params.File, params.FileHeader, err = c.Request.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid or missing params: media", "error": err.Error()})
			return
		}

		params.Content = c.PostForm("content")
		userID := mid.UserID(c)

		article, err := sr.Create(userID, params)
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
	sr := services.NewArticleService(ar.a)

	return func(c *gin.Context) {
		queries, err := validateQueries(
			c.Query("limit"),
			c.Query("offset"),
			c.Query("sort"),
			c.Query("order"),
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "missing or invalid params",
				"error":   err.Error(),
			})
			return
		}

		articles, err := sr.ListAll(queries)

		c.JSON(http.StatusOK, articles)
	}
}

// validateQueries validating query params

/*validateQueries
validating query params
return limit, offset, sort, order and err
*/
func validateQueries(args ...string) (*services.ArticleListQueryParams, error) {
	limit := Limit(args[0])
	if err := limit.Valid(); err != nil {
		return nil, err
	}

	offset := Offset(args[1])
	if err := offset.Valid(); err != nil {
		return nil, err
	}

	sort := Sort(args[2])
	if err := sort.Valid(); err != nil {
		return nil, err
	}

	order := Order(args[3])
	if err := order.Valid(); err != nil {
		return nil, err
	}

	return &services.ArticleListQueryParams{
		sort.String(),
		order.String(),
		limit.Int(),
		offset.Int()}, nil
}
