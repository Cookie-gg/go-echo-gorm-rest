package handler

import (
	"go-echo-gorm-rest/model"
	"go-echo-gorm-rest/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService    *service.UserService
	profileService *service.ProfileService
}

func CreateUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{service.CreateUserService(db), service.CreateProfileService(db)}
}

// CreateUser godoc
// @Tags          User
// @Summary       create a user by id
// @Description   create a user by id
// @Accept        json
// @Produce       json
// @Param    			body body model.User  true "request body"
// @Success       200 {object} model.User
// @Failure       500
// @Router        /user [post]
func (h *UserHandler) CreateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := h.userService.Create(&user); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, user)
}

// GetUser godoc
// @Tags          User
// @Summary       get user by id
// @Description   get user by id
// @Produce       json
// @Param         id path int true "User ID"
// @Success       201 {object} model.User
// @Failure       500
// @Router        /user/{id} [get]
func (h *UserHandler) GetUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := h.userService.Get(&user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// GetUserProfile godoc
// @Tags          User
// @Summary       get a user profile by user id
// @Description   get a user profile by user id
// @Produce       json
// @Param         id path int true "User ID"
// @Success       200 {object} model.User
// @Failure       500
// @Router        /user/{id}/profile [get]
func (h *UserHandler) GetUserProfile(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	profile := model.Profile{UserID: user.ID}
	if err := h.profileService.Get(&profile); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, profile)
}

// UpdateUser godoc
// @Tags          User
// @Summary       update a user
// @Description   update a user
// @Accept        json
// @Param    			body body model.User  true "request body"
// @Success       204
// @Failure       500
// @Router        /user [put]
func (h *UserHandler) UpdateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := h.userService.Update(&user); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

// UpdateUserProfile godoc
// @Tags          User
// @Summary       update a profile by user id
// @Description   update a profile by user id
// @Accept        json
// @Param    			body body model.Profile  true "request body"
// @Param         id path int true "User ID"
// @Success       204
// @Failure       500
// @Router        /user/{id}/profile [put]
func (h *UserHandler) UpdateUserProfile(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	profile := model.Profile{UserID: user.ID}
	if err := h.profileService.Update(&profile); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

// DeleteUser godoc
// @Tags          User
// @Summary       delete a user by user id
// @Description   delete a user by user id
// @Param         id path int true "User ID"
// @Success       204
// @Failure       500
// @Router        /user/{id} [delete]
func (h *UserHandler) DeleteUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := h.userService.Delete(&user); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
