package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetBooksController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, GetBooksController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetBookController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, GetBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreateBookController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, CreateBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteBookController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/books/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, DeleteBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateBookController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/books/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, UpdateBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
