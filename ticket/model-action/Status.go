package modelaction

import (
	"ticket/db"
	"ticket/model"
)

//GetAllStatus - Insert New data
func GetAllStatus(status *[]model.Status) (err error) {
	if err = db.GetDB().Find(&status).Error; err != nil {
		return err
	}
	return nil
}

//GetTicketByStatusID - Fetch all ticket by StatusID
func GetTicketByStatusID(ticket *[]model.Ticket, id string) (err error) {
	if err = db.GetDB().Where("status_id = ?", id).Preload("Status").Preload("Contact").Order("updated_at DESC").Find(&ticket).Error; err != nil {
		return err
	}
	return nil
}
