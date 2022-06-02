package modals

import (
	db "restapi-tugas/database"
)

type User struct {
	Id       uint64 `json:"UserId" sql:"AUTO_INCREMENT" `
	Name     string `json:"UserName"`
	Password string `json:"UserPassword" `
	Email    string `json:"UserEmail" binding:"required"`
}

// FindAllUsers
func FindAllUsers() []User {
	users := []User{}
	db.DBconnect.Find(&users)
	return users
}

// FindByUserId
func FindByUserId(userId string) User {
	user := User{}
	db.DBconnect.Where("id = ?", userId).First(&user)
	return user
}

// CreateUser
func CreateUser(user User) User {
	db.DBconnect.Create(&user)
	return user
}

// DeleteUser
func DeleteUser(userId string) bool {
	user := User{}
	result := db.DBconnect.Where("id = ?", userId).Delete(&user)
	return result.RowsAffected > 0
}

// UpdateUser
func UpdateUser(userId string, user User) User {
	db.DBconnect.Model(&user).Where("id = ?", userId).Updates(user)
	return user
}

// Check if user exists
func LoginUser(name string, password string) User {
	user := User{}
	db.DBconnect.Where("name = ? and password = ?", name, password).First(&user)
	return user
}
