package main

import (
	"go-echo-gorm-rest/model"
	"go-echo-gorm-rest/router"
)

func main() {
	gormDB, _ := model.InitDB()
	db, _ := gormDB.DB()
	defer db.Close()

	e := router.NewRouter(gormDB)
	e.Logger.Fatal(e.Start(":3000"))
}
