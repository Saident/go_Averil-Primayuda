package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	for _, user := range users {
		if user.Id == id {
			u := user
			return c.JSON(http.StatusOK, u)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "User not found",
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]string{
				"message": "User deleted successfully",
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "User not found",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	var updateUser User

	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	if err := c.Bind(&updateUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to parse request body",
		})
	}

	for i, user := range users {
		if user.Id == id {
			users[i] = updateUser
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "User updated successfully",
				"user":    updateUser,
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "User not found",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
