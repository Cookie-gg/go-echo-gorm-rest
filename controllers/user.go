package controllers

import (
	"github.com/labstack/echo/v4"
	"go-echo-gorm-rest/models"
	"net/http"
)

func CreateUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	models.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	models.DB.First(&user, id)
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	models.DB.Save(&user)
	return c.NoContent(http.StatusNoContent)
}

func DeleteUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	models.DB.Delete(&user)
	return c.NoContent(http.StatusNoContent)
}
