package models

import "github.com/google/uuid"

// Comment ...
type Comment struct {
	BaseModel
	ArticleID uuid.UUID `json:"articleId" gorm:"type:uuid;column:article_foreign_key;not null;"`
	UserID    uuid.UUID `json:"userId" gorm:"type:uuid;column:user_foreign_key;not null;"`
	Text      string    `json:"text" gorm:"type:text;not null"`
}
