package models

import (
	"github.com/cayohollanda/gorm-crud/cmd/config"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllPersons(p *[]Person) (err error) {
	if err = config.DB.Find(p).Error; err != nil {
		return err
	}
	return nil
}

func FindPerson(p *Person, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(p).Error; err != nil {
		return err
	}
	return nil
}

func AddPerson(p *Person) (err error) {
	if err := config.DB.Create(p).Error; err != nil {
		return err
	}
	return nil
}

func DeletePerson(p *Person, id string) (err error) {
	if err := config.DB.Where("id = ?", id).Delete(p).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePerson(p *Person) (err error) {
	if err := config.DB.Save(p).Error; err != nil {
		return err
	}
	return nil
}
