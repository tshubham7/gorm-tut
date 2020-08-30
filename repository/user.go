package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/tshubham7/gorm-articles/models"
)

type user struct {
	db *gorm.DB
}

// UserRepo ..
type UserRepo interface {
	// create new user
	Create(user *models.User) *gorm.DB

	// get user by id
	GetByID(id string) (*models.User, error)

	// get user by email
	GetByEmail(id string) (*models.User, error)
}

// NewUserRepo ...
func NewUserRepo(db *gorm.DB) UserRepo {
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

// GetByEmail ...
func (u *user) GetByEmail(email string) (*models.User, error) {
	var user models.User

	err := u.db.Table("users").Where("email=?", email).First(&user).Error
	return &user, err
}
