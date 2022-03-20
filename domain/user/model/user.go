package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                string
	RoleID            string
	Name              string
	Email             string
	EncryptedPassword string
	Token             string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New().String()
	return
}
