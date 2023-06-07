package controller

import (
	"net/http"
	"strconv"

	"praktikum_22/config"
	"praktikum_22/helper"
	"praktikum_22/model"

	"github.com/labstack/echo"
)

// get all items
func GetItemsController(c echo.Context) error {
	var items []model.Item
	keyword := c.QueryParam("keyword")

	if err := config.DB.Where("name LIKE ?", "%"+keyword+"%").Find(&items).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all items",
		"items":   items,
	})
}

// get item by id
func GetItemController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	var items model.Item
	if err := config.DB.First(&items, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get item by id",
		"item":    items,
	})
}

func GetItemByCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	var items model.Item
	if err := config.DB.Where("category = ?", id).Find(&items).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get item by category",
		"item":    items,
	})
}

// create new item
func CreateItemController(c echo.Context) error {
	uuid, _ := helper.NewGoogleUUID().GenerateUUID()
	item := model.Item{}
	item.ID = uuid
	c.Bind(&item)

	if err := config.DB.Save(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new item",
		"item":    item,
	})
}

// delete item by id
func DeleteItemController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	var item model.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete item by id",
	})
}

// update item by id
func UpdateItemController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	var item model.Item

	if err := config.DB.First(&item, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	c.Bind(&item)

	if err := config.DB.Save(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update item by id",
		"item":    item,
	})
}
