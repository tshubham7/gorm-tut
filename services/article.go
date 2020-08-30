package services

import (
	"mime/multipart"

	"github.com/tshubham7/gorm-articles/models"
	"github.com/tshubham7/gorm-articles/repository"
	"github.com/tshubham7/gorm-articles/tools/image"
)

type article struct {
	a repository.ArticleRepo
}

// ArticleCreateRequest ...
type ArticleCreateRequest struct {
	Content    string `json:"content"`
	File       multipart.File
	FileHeader *multipart.FileHeader
}

// ArticleListQueryParams ...
type ArticleListQueryParams struct {
	Sort   string
	Order  string
	Limit  int32
	Offset int32
}

// ToModel ...
func (a ArticleCreateRequest) ToModel(userID string) *models.Article {

	// process the image and get the image id
	imageID, _ := image.HandleFile(a.File, a.FileHeader)

	return &models.Article{
		Content: a.Content,
		Image:   imageID,
		UserID:  userID,
	}
}

// ArticleService ...
type ArticleService interface {
	// create new article
	Create(currentUserID string, Request ArticleCreateRequest) (*models.Article, error)

	// list all articles
	ListAll(queries *ArticleListQueryParams) ([]models.Article, error)
}

// NewArticleService ...
func NewArticleService(a repository.ArticleRepo) ArticleService {
	return &article{a}
}

// Create ...
func (ar article) Create(userID string, article ArticleCreateRequest) (*models.Article, error) {

	a := article.ToModel(userID)
	err := ar.a.Create(a)

	return a, err
}

// ListAll ...
func (ar article) ListAll(q *ArticleListQueryParams) ([]models.Article, error) {
	return ar.a.ListAll(q.Sort, q.Order, q.Limit, q.Offset)
}
