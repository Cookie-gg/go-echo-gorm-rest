package main

import (
	_ "go-echo-gorm-rest/docs"

	"go-echo-gorm-rest/model"
	"go-echo-gorm-rest/router"
)

// @title         go-echo-gorm-rest
// @version       1.0
// @license.name  cookie_gg
// @host localhost:3000
// @BasePath /
func main() {
	gormDB, _ := model.InitDB()
	db, _ := gormDB.DB()
	defer db.Close()

	e := router.NewRouter(gormDB)
	e.Logger.Fatal(e.Start(":3000"))
}
