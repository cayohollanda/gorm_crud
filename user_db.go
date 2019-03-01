package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func createUserRepo(newUser User) {
	getConnection().Create(&newUser)
	getConnection().Close()
}

func deleteUserRepo(ID int) {
	getConnection().Where("id = ?", ID).Delete(&User{})
	getConnection().Close()
}

func getUsersListRepo() []User {
	var users []User
	getConnection().Find(&users)
	getConnection().Close()

	return users
}

func getUserRepo(ID int) User {
	var user User
	getConnection().Find(&user, ID)
	getConnection().Close()

	return user
}
