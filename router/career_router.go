package router

import (
	"chat_upgrade/controller"

	"github.com/labstack/echo/v4"
)

func RegisterCareerRoutes(e *echo.Group, careerController controller.ICareerController) {
	e.POST("/careers", careerController.CreateCareer)
	e.GET("/careers", careerController.GetCareers)
	e.DELETE("/careers/:id", careerController.DeleteCareer)
}
