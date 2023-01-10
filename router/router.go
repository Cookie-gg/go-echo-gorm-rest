package router

import (
	"go-echo-gorm-rest/handler"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *echo.Echo {
	e := echo.New()

	h := handler.CreateUserHandler(db)

	e.GET("/user/:id", h.GetUser)
	e.GET("/user/:id/profile", h.GetUserProfile)
	e.POST("/user", h.CreateUser)
	e.PUT("/user/:id", h.UpdateUser)
	e.PUT("/user/:id/profile", h.UpdateUserProfile)
	e.DELETE("/user/:id", h.DeleteUser)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return e
}
