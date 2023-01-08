package handler_test

import (
	"go-echo-gorm-rest/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Setup(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, echo.Context, *httptest.ResponseRecorder, *http.Request) {
	mockDB, sqlMock, err := model.InitDBMock()
	if err != nil {
		t.Fatal(err)
	}

	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	c := e.NewContext(req, rec)

	return mockDB, sqlMock, c, rec, req
}
