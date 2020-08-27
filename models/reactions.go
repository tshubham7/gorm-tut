package models

import "github.com/google/uuid"

// Comment ...
type Comment struct {
	BaseModel
	ArticleID uuid.UUID `gorm:"type:uuid;column:article_foreign_key;not null;"`
	UserID    uuid.UUID `gorm:"type:uuid;column:user_foreign_key;not null;"`
	Text      string    `gorm:"type:text;not null"`
}
