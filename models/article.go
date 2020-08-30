package models

// Article ...
type Article struct {
	BaseModel
	UserID  string `json:"userId" gorm:"type:uuid;column:user_id;not null;"`
	Content string `json:"content" gorm:"type:text;not null"`
	Image   string `json:"image" gorm:"type:text;not null"`
}
