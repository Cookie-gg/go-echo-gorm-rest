package service_test

import (
	"go-echo-gorm-rest/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

func Setup(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	mockDB, sqlMock, err := model.InitDBMock()
	if err != nil {
		t.Fatal(err)
	}

	return mockDB, sqlMock
}
