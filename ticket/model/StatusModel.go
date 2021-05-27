package model

// Status - model
type Status struct {
	ID     uint64 `gorm:"primary_key auto_increment" json:"id" `
	Status string `json:"status" gorm:"type:varchar(50)"`
}
