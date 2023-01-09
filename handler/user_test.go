package handler_test

import (
	"go-echo-gorm-rest/handler"
	"go-echo-gorm-rest/mock"
	"go-echo-gorm-rest/model"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"
)

var (
	mockDB  *gorm.DB
	sqlMock sqlmock.Sqlmock
	e       *echo.Echo
	err     error
)

func init() {
	mockDB, sqlMock, err = model.InitDBMock()
	if err != nil {
		log.Fatal(err)
	}
	e = echo.New()
}

func TestCreateUser(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(mock.UserJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	c := e.NewContext(req, rec)

	h := handler.CreateUserHandler(mockDB)

	query := "INSERT INTO `users`"

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(regexp.QuoteMeta(query)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	if assert.NoError(t, h.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, mock.UserJson, rec.Body.String())
	}
}

func TestGetUser(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	c := e.NewContext(req, rec)

	h := handler.CreateUserHandler(mockDB)

	c.SetPath("/user/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	query := "SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = ? LIMIT 1"

	sqlMock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(sqlmock.NewRows(model.UserColumns).AddRow(mock.User.ID, mock.User.Name, mock.User.CreatedAt, mock.User.UpdatedAt, mock.User.DeletedAt))

	if assert.NoError(t, h.GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mock.UserJson, rec.Body.String())
	}
}
