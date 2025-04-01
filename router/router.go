package router

import (
	"chat_upgrade/controller"
	"chat_upgrade/middleware"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)


func NewRouter(
	uc controller.IUserController,
	ec controller.IExperienceController,
	cc controller.ICareerController,
) *echo.Echo {
	e := echo.New()

	// CORS 設定
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{
			echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken, echo.HeaderAuthorization,
		},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	// CSRF 保護
	e.Use(echoMiddleware.CSRFWithConfig(echoMiddleware.CSRFConfig{
		CookiePath:     "/",
		// CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieDomain: "",
		CookieHTTPOnly: true,
		CookieSecure:   false,
		CookieSameSite: http.SameSiteDefaultMode,
		CookieMaxAge:   3600,
	}))

	// ユーザー関連のエンドポイント
	e.POST("/login", uc.Login)
	e.GET("/csrf", uc.CsrfToken)

	// 認証が必要なエンドポイント
	authGroup := e.Group("") // 認証が必要なルートグループ
	authGroup.Use(middleware.JWTMiddleware) // JWT Middleware を適用

	authGroup.GET("/getMe", func(c echo.Context) error {
		fmt.Println("認証が必要なエンドポイントに到達")
		return uc.GetMe(c)
	})

	authGroup.POST("/uploadmyicon", func(c echo.Context) error {
		fmt.Println("認証済みエンドポイント `/uploadmyicon` に到達")
		return uc.UploadUserIcon(c)
	})

	// ✅ experience ルートの追加
	NewExperienceRouter(ec, authGroup)

	// ✅ career ルートの登録
	RegisterCareerRoutes(authGroup, cc)

	return e
}
