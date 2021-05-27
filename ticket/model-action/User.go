package modelaction

import (
	"fmt"
	"ticket/db"
	"ticket/model"
)

//GetAllUsers - Fetch all user data
func GetAllUsers(user *[]model.User) (err error) {
	if err = db.GetDB().Find(&user).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser - Insert New data
func CreateUser(user *model.User) (err error) {
	if err = db.GetDB().Create(&user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID - Fetch only one user by Id
func GetUserByID(user *model.User, id string) (err error) {
	if err = db.GetDB().Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser - Update user
func UpdateUser(user *model.User, id string) (err error) {
	fmt.Println(user)
	db.GetDB().Save(&user)
	return nil
}

//DeleteUser - Delete user
func DeleteUser(user *model.User, id string) (err error) {
	db.GetDB().Where("id = ?", id).Delete(&user)
	return nil
}

//CheckHasUser - Fetch Username
func CheckHasUser(user *model.User, username string) (err error) {
	if err = db.GetDB().First(&user, "username = ?", username).Error; err != nil {
		return err
	}
	return nil
}
