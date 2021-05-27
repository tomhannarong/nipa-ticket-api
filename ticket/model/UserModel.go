package model

import (
	"time"

	"gorm.io/gorm"
)

// User - Model
type User struct {
	gorm.Model
	ID        uint64 `gorm:"primary_key auto_increment" json:"id" `
	Username  string `gorm:"unique" form:"username" binding:"required"`
	Password  string `form:"password" binding:"required"`
	Level     string `gorm:"default:normal"`
	StatusID  uint64 `json:"-"`
	Status    Status `json:"status" gorm:"foreignkey:StatusID"`
	CreatedAt time.Time
}
