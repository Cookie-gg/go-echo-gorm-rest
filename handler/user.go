package handler

import (
	"go-echo-gorm-rest/model"
	"go-echo-gorm-rest/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService *service.UserService
}

func CreateUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{service.CreateUserService(db)}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	err := h.userService.Create(&user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	err := h.userService.Get(&user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	h.userService.Update(&user)
	return c.NoContent(http.StatusNoContent)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	h.userService.Delete(&user)
	return c.NoContent(http.StatusNoContent)
}
