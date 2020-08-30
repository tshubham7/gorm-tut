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

	// delete article
	Delete() gin.HandlerFunc

	// article detail
	Detail() gin.HandlerFunc
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
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to fetch articles",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, articles)
	}
}

// Delete ...
func (ar article) Delete() gin.HandlerFunc {
	sr := services.NewArticleService(ar.a)

	return func(c *gin.Context) {
		id := c.Param("id")

		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "missing or invalid params: id ",
				"error":   "can not allow blank article id",
			})
			return
		}

		err := sr.Delete(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to delete article",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
	}
}

// Detail ...
func (ar article) Detail() gin.HandlerFunc {
	sr := services.NewArticleService(ar.a)

	return func(c *gin.Context) {
		id := c.Param("id")

		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "missing or invalid params: id ",
				"error":   "can not allow blank article id",
			})
			return
		}

		a, err := sr.Detail(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to delete article",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, a)
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
		Sort:   sort.String(),
		Order:  order.String(),
		Limit:  limit.Int(),
		Offset: offset.Int()}, nil
}
