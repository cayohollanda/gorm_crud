package main

import "github.com/jinzhu/gorm"

func getConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/gorm_crud?charset=utf8&parseTime=True")
	checkErr(err)
	return db
}

func loadDb() {
	getConnection().AutoMigrate(&User{})
	getConnection().Close()
}
