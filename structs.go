package main

import "time"

// User struct
type User struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string `gorm:"type:varchar(50);not null" json:"name"`
	Username  string `gorm:"type:varchar(20);not null; unique" json:"username"`
	Email     string `gorm:"type:varchar(50);not null; unique" json:"email"`
	Role      string `gorm:"type:varchar(10);not null;default:\"USER\"" json:"role"`
}

// Error struct
type Error struct {
	message string
}
