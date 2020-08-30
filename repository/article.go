package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/tshubham7/gorm-articles/models"
)

type article struct {
	db *gorm.DB
}

// ArticleRepo ...
type ArticleRepo interface {
	// create a new article
	Create(article *models.Article) error

	// list all articles
	ListAll(sort, order string, limit, offset int32) ([]models.Article, error)

	// list articles by user id
	ListByUser(userID string) ([]models.Article, error)
}

// NewArticleRepo ...
func NewArticleRepo(db *gorm.DB) ArticleRepo {
	return &article{db}
}

// Create ...
func (a *article) Create(article *models.Article) error {
	return a.db.Create(article).Error
}

// ListByUser ...
func (a *article) ListByUser(userID string) ([]models.Article, error) {
	var articles = []models.Article{}
	result := a.db.Table("users").Where("user_id=?", userID).First(&articles)
	return articles, result.Error
}

// ListAll ...
func (a *article) ListAll(sort, order string, limit, offset int32) ([]models.Article, error) {
	var articles = []models.Article{}

	result := a.db.Table("articles").Order(fmt.Sprintf("%s %s", sort, order))
	result = result.Limit(limit).Offset(offset).Find(&articles)

	return articles, result.Error
}
