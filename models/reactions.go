package models

// Comment ...
type Comment struct {
	BaseModel
	ArticleID string `json:"articleId" gorm:"type:uuid;column:article_id;not null;"`
	UserID    string `json:"userId" gorm:"type:uuid;column:user_id;not null;"`
	Text      string `json:"text" gorm:"type:text;not null"`
}
