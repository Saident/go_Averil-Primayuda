package route

import (
	"praktikum_22/constants"
	"praktikum_22/controller"
	m "praktikum_22/middleware"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/users", controller.CreateUserController)
	e.POST("/login", controller.LoginUserController)

	eJwt := e.Group("")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))

	eJwt.GET("/items", controller.GetItemsController)
	eJwt.GET("/items/:id", controller.GetItemController)
	eJwt.GET("/items/category/:id", controller.GetItemByCategoryController)
	eJwt.POST("/items", controller.CreateItemController)
	eJwt.DELETE("/items/:id", controller.DeleteItemController)
	eJwt.PUT("/items/:id", controller.UpdateItemController)

	m.LogMiddleware(e)
	return e
}
