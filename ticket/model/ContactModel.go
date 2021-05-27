package model

// Contact - model
type Contact struct {
	ID       uint64 `gorm:"primary_key auto_increment" json:"id" `
	Name     string `json:"name" gorm:"type:varchar(256)"`
	TicketID uint64 `json:"-"`
	Tel      string `json:"tel" gorm:"type:varchar(10)"`
}
