package models

import (
	"github.com/cayohollanda/gorm-crud/cmd/config"
	"github.com/cayohollanda/gorm-crud/cmd/utils"
)

func AddUser(u *User) (err error) {
	var hash utils.Hash
	if u.Password, err = hash.Generate(u.Password); err != nil {
		return err
	}
	if err := config.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func FindUserByUsername(u *User) (err error) {
	if err := config.DB.Where("username = ?", u.Username).First(u).Error; err != nil {
		return err
	}
	return nil
}
