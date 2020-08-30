package services

import (
	"github.com/tshubham7/gorm-articles/models"
	"github.com/tshubham7/gorm-articles/repository"
)

type article struct {
	a repository.ArticleRepo
}

// ArticleCreateRequest ...
type ArticleCreateRequest struct {
	Content string `json:"content"`
	Image   string `json:"image"`
}

// ToModel ...
func (a ArticleCreateRequest) ToModel(userID string) *models.Article {
	return &models.Article{
		Content: a.Content,
		Image:   a.Image,
		UserID:  userID,
	}
}

// ArticleService ...
type ArticleService interface {
	Create(ArticleCreateRequest) (*models.Article, error)
}

// NewArticleService ...
func NewArticleService(a repository.ArticleRepo) ArticleService {
	return &article{a}
}

// Create ...
func (ar article) Create(article ArticleCreateRequest) (*models.Article, error) {

	a := article.ToModel("userid")
	err := ar.a.Create(a)

	return a, err
}
