package seed

import (
	"ticket/model"

	"gorm.io/gorm"
)

var status = []model.Status{
	model.Status{
		Status: "pending",
	},
	model.Status{
		Status: "accepted",
	},
	model.Status{
		Status: "resolved",
	},
	model.Status{
		Status: "rejected",
	},
}

// Load - data status ticket
func Load(db *gorm.DB) {
	var statusQuery model.Status
	err := db.First(&statusQuery).Error
	if err == gorm.ErrRecordNotFound {
		for i, _ := range status {
			db.Model(&model.Status{}).Create(&status[i])
		}
	}
}
