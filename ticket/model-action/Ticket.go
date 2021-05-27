package modelaction

import (
	"fmt"
	"ticket/db"
	"ticket/model"
)

//GetAllTickets - Fetch all Ticket data
func GetAllTickets(ticket *[]model.Ticket) (err error) {
	if err = db.GetDB().Preload("Status").Preload("Contact").Order("updated_at DESC").Find(&ticket).Error; err != nil {
		return err
	}
	return nil
}

//GetFilterTickets - Fetch by Filter Ticket data
func GetFilterTickets(ticket *[]model.Ticket, StatusID string) (err error) {
	if err = db.GetDB().Preload("Status").Preload("Contact").Where("status_id = ?", StatusID).Order("updated_at DESC").Find(&ticket).Error; err != nil {
		return err
	}
	return nil
}

//CreateTicket - Insert New data
func CreateTicket(ticket *model.Ticket) (err error) {
	if err = db.GetDB().Create(&ticket).Error; err != nil {
		return err
	}
	return nil
}

//GetTicketByID - Fetch only one Ticket by Id
func GetTicketByID(ticket *model.Ticket, id string) (err error) {
	if err = db.GetDB().Preload("Status").Preload("Contact").Where("id = ?", id).First(&ticket).Error; err != nil {
		return err
	}
	return nil
}

//UpdateTicket - Update Ticket
func UpdateTicket(ticket *model.Ticket, id string) (err error) {
	fmt.Println(ticket)
	db.GetDB().Updates(&ticket)
	return nil
}

//DeleteTicket - Delete Ticket
func DeleteTicket(ticket *model.Ticket, id string) (err error) {
	db.GetDB().Where("id = ?", id).Delete(&ticket)
	return nil
}
