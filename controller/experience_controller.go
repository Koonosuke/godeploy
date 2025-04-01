package controller

import (
	"chat_upgrade/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IExperienceController interface {
	CreateExperience(c echo.Context) error
	GetAllExperiences(c echo.Context) error
	DeleteExperience(c echo.Context) error
}

type experienceController struct {
	eu usecase.IExperienceUsecase
}

func NewExperienceController(eu usecase.IExperienceUsecase) IExperienceController {
	return &experienceController{eu}
}

func (ec *experienceController) CreateExperience(c echo.Context) error {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "認証が必要です",
		})
	}

	// フォームデータを取得
	title := c.FormValue("title")
	techStack := c.FormValue("tech_stack")
	content := c.FormValue("content")
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "アイコン画像が提供されていません",
		})
	}

	// 経験情報の作成
	experience, err := ec.eu.CreateExperience(userID, title, techStack, content, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "経験情報の作成に失敗しました",
		})
	}

	return c.JSON(http.StatusOK, experience)
}

func (ec *experienceController) GetAllExperiences(c echo.Context) error {
	experiences, err := ec.eu.GetAllExperiences()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "経験情報の取得に失敗しました",
		})
	}
	return c.JSON(http.StatusOK, experiences)
}

func (ec *experienceController) DeleteExperience(c echo.Context) error {
	// 文字列IDを取得し、整数に変換
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32) // 10進数でパース
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "IDの形式が正しくありません",
		})
	}

	// 削除処理
	if err := ec.eu.DeleteExperience(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "経験情報の削除に失敗しました",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "経験情報が削除されました",
	})
}
