package controller

import (
	"chat_upgrade/model"
	"chat_upgrade/usecase"
	"fmt"
	"net/http"
	"os"

	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	Login(c echo.Context) error
	GetMe(c echo.Context) error
	UploadUserIcon(c echo.Context) error
    CsrfToken(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

// ログイン処理
func (uc *userController) Login(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "リクエストデータの形式が正しくありません",
		})
	}

	if user.Username == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "username と password は必須です",
		})
	}

	token, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "認証に失敗しました",
		})
	}
// cookie.Secure = true の意味
//このCookieは HTTPS 通信のときだけ送信される という設定！
	// トークンをCookieに設定
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	
	if os.Getenv("ENV") == "production" {
		cookie.Domain = os.Getenv("API_DOMAIN")     // e.g. nextdeploy-navy.vercel.app
		cookie.Secure = true                         // HTTPS 通信に限定
		cookie.SameSite = http.SameSiteNoneMode      // クロスサイト Cookie 対応
	} else {
		cookie.Secure = false
		cookie.SameSite = http.SameSiteLaxMode
	}
	
	cookie.HttpOnly = true
	c.SetCookie(cookie)
	

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

// ユーザー情報を取得 (`/me`)
func (uc *userController) GetMe(c echo.Context) error {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "認証が必要です",
		})
	}

	user, err := uc.uu.GetMe(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "ユーザー情報の取得に失敗しました",
		})
	}

	return c.JSON(http.StatusOK, user)
}

// ユーザーアイコンをアップロード
func (uc *userController) UploadUserIcon(c echo.Context) error {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "認証が必要です",
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "ファイルが提供されていません",
		})
	}

	iconURL, err := uc.uu.UploadUserIcon(userID, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "アイコンのアップロードに失敗しました",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":  "アイコンがアップロードされました",
		"icon_url": iconURL,
	})
}

func getUserIDFromContext(c echo.Context) (uint, error) {
    userIDVal := c.Get("user_id")
    if userIDVal == nil {
        fmt.Println("Context に user_id が存在しません")
        return 0, fmt.Errorf("user_id が存在しない")
    }

    userID, ok := userIDVal.(uint)
    if !ok {
        fmt.Println("user_id の型が正しくありません:", userIDVal)
        return 0, fmt.Errorf("user_id の型が正しくありません")
    }

    fmt.Printf("Context から取得した user_id: %d\n", userID)
    return userID, nil
}


func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
