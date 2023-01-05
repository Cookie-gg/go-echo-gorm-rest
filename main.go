package main

import (
	"github.com/labstack/echo/v4"
	"go-echo-gorm-rest/controllers"
	"go-echo-gorm-rest/models"
)

func main() {
	e := echo.New()

	db, _ := models.DB.DB()
	defer db.Close()

	e.GET("/user/:id", controllers.GetUser)
	e.POST("/user", controllers.CreateUser)
	e.PUT("/user/:id", controllers.UpdateUser)
	e.DELETE("/user/:id", controllers.DeleteUser)

	e.Logger.Fatal(e.Start(":3000"))
}
