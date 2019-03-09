package models

import (
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	Name string `json:"name"`
}

func (p *Person) TableName() string {
	return "persons"
}
