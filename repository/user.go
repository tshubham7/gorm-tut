package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/tshubham7/gorm-articles/models"
)

type user struct {
	db *gorm.DB
}

// UserService ..
type UserService interface {
	Create(user *models.User) *gorm.DB
	GetByID(id string) (*models.User, error)
}

// NewUserService ...
func NewUserService(db *gorm.DB) UserService {
	return &user{db}
}

// Create ...
func (u *user) Create(user *models.User) *gorm.DB {
	return u.db.Create(user)
}

// GetByID ...
func (u *user) GetByID(id string) (*models.User, error) {
	var user models.User

	err := u.db.Table("users").Find(&user, id).Error
	return &user, err
}
