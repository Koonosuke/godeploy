package model

import "time"

// User struct represents the user entity in the database

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	UserIcon  string    `gorm:"default:'/default-icon.png'" json:"userIcon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserResponse struct for responding to clients
type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	UserIcon string `json:"userIcon"`
}
