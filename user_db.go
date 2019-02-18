package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func createUserRepo(newUser User) (u User, e Error) {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/gorm_crud?charset=utf8&parseTime=True")
	checkErr(err)

	db.Create(newUser)

	var ret User
	db.First(&ret)

	db.Close()

	return ret, Error{}
}

func getUsersListRepo() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/gorm_crud?charset=utf8&parseTime=True")
	checkErr(err)

	db.Close()
}

func loadDb() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/gorm_crud?charset=utf8&parseTime=True&loc=Local")
	checkErr(err)

	db.AutoMigrate(&User{})
	db.Close()
}
