package modelaction

import (
	"fmt"
	"ticket/db"
	"ticket/model"
)

//CreateContact - Insert New data
func CreateContact(contact *model.Contact) (err error) {
	if err = db.GetDB().Create(&contact).Error; err != nil {
		return err
	}
	return nil
}

//GetContactByID - Fetch only one Contact by Id
func GetContactByID(contact *model.Contact, id string) (err error) {
	if err = db.GetDB().Where("id = ?", id).First(&contact).Error; err != nil {
		return err
	}
	return nil
}

//UpdateContact - Update Contact
func UpdateContact(contact *model.Contact) (err error) {
	fmt.Println(contact)
	db.GetDB().Updates(&contact)
	return nil
}

//DeleteContact - Delete Contact
func DeleteContact(contact *model.Contact, id string) (err error) {
	db.GetDB().Where("id = ?", id).Delete(&contact)
	return nil
}
