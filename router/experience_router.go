package router

import (
	"chat_upgrade/controller"

	"github.com/labstack/echo/v4"
)

// NewExperienceRouter は、/experiences に関するルートを定義する
func NewExperienceRouter(ec controller.IExperienceController, authGroup *echo.Group) {
	authGroup.POST("/experiences", ec.CreateExperience)
	authGroup.GET("/experiences", ec.GetAllExperiences)
	authGroup.DELETE("/experiences/:id", ec.DeleteExperience)
}
