package controller

import (
	"chat_upgrade/model"
	"chat_upgrade/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ICareerController interface {
	CreateCareer(c echo.Context) error
	GetCareers(c echo.Context) error
	DeleteCareer(c echo.Context) error
}

type CareerController struct {
	usecase usecase.ICareerUsecase
}

func NewCareerController(usecase usecase.ICareerUsecase) *CareerController {
	return &CareerController{usecase: usecase}
}

func (cc *CareerController) CreateCareer(c echo.Context) error {
	var career model.Career
	if err := c.Bind(&career); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	if err := cc.usecase.CreateCareer(&career); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create career"})
	}

	return c.JSON(http.StatusCreated, career)
}

func (cc *CareerController) GetCareers(c echo.Context) error {
	careers, err := cc.usecase.GetCareers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch careers"})
	}
	return c.JSON(http.StatusOK, careers)
}

func (cc *CareerController) DeleteCareer(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid career ID"})
	}

	if err := cc.usecase.DeleteCareer(id); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete career"})
	}

	return c.NoContent(http.StatusNoContent)
}
