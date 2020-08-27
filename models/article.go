package models

import "github.com/google/uuid"

// Article ...
type Article struct {
	BaseModel
	UserID  uuid.UUID `gorm:"type:uuid;column:user_foreign_key;not null;"`
	Content string    `gorm:"type:text;not null"`
	Image   string    `gorm:"type:text;not null"`
}
