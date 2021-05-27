package model

import (
	"time"

	"gorm.io/gorm"
)

// Ticket - model
type Ticket struct {
	gorm.Model
	ID          uint64    `gorm:"primary_key auto_increment" json:"id" `
	Title       string    `json:"title" gorm:"type:varchar(100)"`
	Description string    `json:"description" gorm:"type:varchar(256)"`
	StatusID    uint64    `json:"-"`
	Status      Status    `json:"status" binding:"required" gorm:"foreignkey:StatusID"`
	Contact     Contact   `json:"contact" binding:"required" gorm:"foreignkey:TicketID"`
	CreatedAt   time.Time `json:"time"`
}
