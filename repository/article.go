package repository

import (
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

	// list articles by user id
	List(userID string) ([]models.Article, error)
}

// NewArticleRepo ...
func NewArticleRepo(db *gorm.DB) ArticleRepo {
	return &article{db}
}

// Create ...
func (a *article) Create(article *models.Article) error {
	return a.db.Create(article).Error
}

// List ...
func (a *article) List(userID string) ([]models.Article, error) {
	var articles = []models.Article{}
	err := a.db.Table("users").Where("user_id=?", userID).First(&articles).Error
	return articles, err
}
