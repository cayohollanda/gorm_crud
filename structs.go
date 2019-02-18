package main

import (
	"github.com/jinzhu/gorm"
)

// User struct
type User struct {
	// with gorm.Model, automatically creates a ID, CreatedAt, UpdatedAt and DeletedAt
	gorm.Model

	Name     string `gorm:"type:varchar(50);not null" json:"name"`
	Username string `gorm:"type:varchar(20);not null; unique" json:"username"`
	Email    string `gorm:"type:varchar(50);not null; unique" json:"email"`
	Role     string `gorm:"type:varchar(10);not null; unique;default:USER" json:"role"`
}

// Error struct
type Error struct {
	message string
}
