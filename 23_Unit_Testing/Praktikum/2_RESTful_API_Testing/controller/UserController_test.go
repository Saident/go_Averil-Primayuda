package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetUsersController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test case: Get all users
	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// TODO: Add more assertions for the response body
	}
}

func TestGetUserController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test case: Get all users
	if assert.NoError(t, GetUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// TODO: Add more assertions for the response body
	}
}

func TestCreateUserController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test case: Get all users
	if assert.NoError(t, CreateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// TODO: Add more assertions for the response body
	}
}

func TestDeleteUserController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, DeleteUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

	}
}

func TestUpdateUserController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test case: Get all users
	if assert.NoError(t, UpdateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// TODO: Add more assertions for the response body
	}
}

func TestLoginUserController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test case: Get all users
	if assert.NoError(t, LoginUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// TODO: Add more assertions for the response body
	}
}
