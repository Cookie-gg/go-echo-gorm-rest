package service_test

import (
	"go-echo-gorm-rest/model"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

var (
	mockDB  *gorm.DB
	sqlMock sqlmock.Sqlmock
	err     error
)

func init() {
	mockDB, sqlMock, err = model.InitDBMock()
	if err != nil {
		log.Fatal(err)
	}
}
