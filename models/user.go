package models

// User ...
type User struct {
	BaseModel
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"type:varchar(100);not null;unique_index"`
	Password string `json:"-" gorm:"not null"`
}
