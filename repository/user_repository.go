package repository

import (
	"chat_upgrade/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByUsername(user *model.User, username string) error
	GetUserByID(user *model.User, userID uint) error // GetUserByID をインターフェースに追加
	UpdateUserIcon(userID uint, iconURL string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

// username でユーザーを取得
func (ur *userRepository) GetUserByUsername(user *model.User, username string) error {
	return ur.db.Where("username = ?", username).First(user).Error
}

// userID でユーザーを取得
func (ur *userRepository) GetUserByID(user *model.User, userID uint) error {
	return ur.db.Where("id = ?", userID).First(user).Error
}

// userID を使ってアイコンの URL を更新
func (ur *userRepository) UpdateUserIcon(userID uint, iconURL string) error {
	return ur.db.Model(&model.User{}).Where("id = ?", userID).Update("user_icon", iconURL).Error
}
