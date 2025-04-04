package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// ✅ Authorization ヘッダーから取得
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			fmt.Println("Authorization ヘッダーがありません")
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "Authorization ヘッダーが必要です",
			})
		}

		// ✅ Bearer トークン形式を確認
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "Bearer トークン形式が必要です",
			})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// ✅ JWT を解析
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("不正な署名アルゴリズム")
			}
			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil || !token.Valid {
			fmt.Println("トークン解析失敗:", err)
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "無効なトークンです",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["user_id"] == nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "ユーザー情報がトークンに含まれていません",
			})
		}

		userID := uint(claims["user_id"].(float64))
		fmt.Println("認証された user_id:", userID)
		c.Set("user_id", userID)

		return next(c)
	}
}


// package middleware

// import (
// 	"fmt"
// 	"net/http"
// 	"os"

// 	"github.com/golang-jwt/jwt/v4"
// 	"github.com/labstack/echo/v4"
// )

// func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// ✅ Cookie からトークンを取得
// 		cookie, err := c.Cookie("token")
// 		if err != nil {
// 			fmt.Println("Cookie 'token' が存在しません:", err)
// 			return c.JSON(http.StatusUnauthorized, echo.Map{
// 				"error": "トークンが提供されていません（Cookie）",
// 			})
// 		}

// 		tokenStr := cookie.Value

// 		// ✅ トークンを解析
// 		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
// 			// 署名アルゴリズムの検証
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("不正な署名アルゴリズムです")
// 			}
// 			return []byte(os.Getenv("SECRET")), nil
// 		})

// 		if err != nil || !token.Valid {
// 			fmt.Println("トークンの解析に失敗:", err)
// 			return c.JSON(http.StatusUnauthorized, echo.Map{
// 				"error": "トークンが無効です",
// 			})
// 		}

// 		// ✅ クレームから user_id を取り出す
// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		if !ok || claims["user_id"] == nil {
// 			fmt.Println("トークンに user_id が含まれていません")
// 			return c.JSON(http.StatusUnauthorized, echo.Map{
// 				"error": "ユーザー情報が含まれていません",
// 			})
// 		}

// 		userID := uint(claims["user_id"].(float64))
// 		fmt.Println("認証された user_id:", userID)

// 		// ✅ Context に user_id を設定
// 		c.Set("user_id", userID)

// 		return next(c)
// 	}
// }
