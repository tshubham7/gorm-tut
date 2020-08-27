package models

// User ...
type User struct {
	BaseModel
	Name     string `gorm:"not null"`
	Email    string `gorm:"type:varchar(100);not null;unique_index"`
	Password string `gorm:"not null"`
}
