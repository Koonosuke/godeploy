package usecase

import (
	"chat_upgrade/model"
	"chat_upgrade/repository"
	"chat_upgrade/validator"
	"mime/multipart"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	Login(user model.User) (string, error)
	GetMe(userID uint) (model.UserResponse, error)
	UploadUserIcon(userID uint, file *multipart.FileHeader) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
	s3 repository.IS3Repository
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator, s3 repository.IS3Repository) IUserUsecase {
	return &userUsecase{ur, uv, s3}
}

// ログイン処理
func (uu *userUsecase) Login(user model.User) (string, error) {
	// ユーザーバリデーション
	if err := uu.uv.ValidateUser(user); err != nil {
		return "", err
	}

	// ユーザー情報の取得
	storedUser := model.User{}
	if err := uu.ur.GetUserByUsername(&storedUser, user.Username); err != nil {
		return "", err
	}

	// パスワードの検証
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return "", err
	}

	// JWTトークンを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	// トークンに署名
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ユーザー情報を取得
func (uu *userUsecase) GetMe(userID uint) (model.UserResponse, error) {
	user := model.User{}
	// userID でユーザーを取得
	if err := uu.ur.GetUserByID(&user, userID); err != nil {
		return model.UserResponse{}, err
	}

	return model.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		UserIcon: user.UserIcon,
	}, nil
}

// ユーザーアイコンをアップロード
func (uu *userUsecase) UploadUserIcon(userID uint, file *multipart.FileHeader) (string, error) {
	// アイコンをS3にアップロード
	iconURL, err := uu.s3.UploadFile(file)
	if err != nil {
		return "", err
	}

	// データベースを更新
	if err := uu.ur.UpdateUserIcon(userID, iconURL); err != nil {
		return "", err
	}

	return iconURL, nil
}
