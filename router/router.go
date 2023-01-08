package router

import (
	"go-echo-gorm-rest/handler"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *echo.Echo {
	e := echo.New()

	h := handler.CreateUserHandler(db)

	e.GET("/user/:id", h.GetUser)
	e.POST("/user", h.CreateUser)
	e.PUT("/user/:id", h.UpdateUser)
	e.DELETE("/user/:id", h.DeleteUser)

	return e
}
