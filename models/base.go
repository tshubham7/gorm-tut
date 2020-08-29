package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// BaseModel contains common columns for all tables.
type BaseModel struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid.String())
}
