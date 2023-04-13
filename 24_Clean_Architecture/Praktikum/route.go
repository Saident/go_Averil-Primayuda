package main

import (
	"clean_architecture/constants"
	"clean_architecture/controller"
	"clean_architecture/repository"
	"clean_architecture/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	jwt "github.com/labstack/echo-jwt"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {

	userRepository := repository.NewUserRepository(db)

	userService := usecase.NewUserUsecase(userRepository)

	userController := controller.NewUserController(userService)

	eJwt := e.Group("")
	eJwt.Use(jwt.JWT([]byte(constants.SECRET_JWT)))
	eJwt.POST("/users", userController.Create)
	eJwt.GET("/users", userController.GetAll)
}
