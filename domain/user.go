package model

import (
	"time"
)

// User ユーザ
type User struct {
	ID        uint      `gorm:"primary_key"`
	UserId    string    `json:""`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
