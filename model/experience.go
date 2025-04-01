package model

import "time"

// Experience struct represents the user experience entity in the database
type Experience struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"` // Foreign key to User
	Title     string    `gorm:"not null" json:"title"`
	TechStack string    `gorm:"not null" json:"tech_stack"` // 技術スタック（カンマ区切り）
	Icon      string    `gorm:"not null" json:"icon"`       // アイコンのパス
	Content   string    `gorm:"type:text" json:"content"`   // 経験の詳細内容
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ExperienceResponse struct for API response
type ExperienceResponse struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Title     string `json:"title"`
	TechStack string `json:"tech_stack"`
	Icon      string `json:"icon"`
	Content   string `json:"content"`
}
