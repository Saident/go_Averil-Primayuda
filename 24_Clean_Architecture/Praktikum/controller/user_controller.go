package controller

import (
	"net/http"

	"clean_architecture/dto"
	"clean_architecture/usecase"

	"github.com/labstack/echo/v4"
)

type UserController interface{}

type userController struct {
	useCase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *userController {
	return &userController{
		userUsecase,
	}
}

func (u *userController) Create(c echo.Context) error {
	var payloads []dto.CreateUserRequest

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	err := u.useCase.Create(payloads)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "Success")
}

func (u *userController) GetAll(c echo.Context) error {
	var payloads []dto.CreateUserRequest

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	user, err := u.useCase.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   user,
	})
}
