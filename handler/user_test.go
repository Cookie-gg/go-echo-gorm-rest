package handler_test

import (
	"go-echo-gorm-rest/handler"
	"go-echo-gorm-rest/mock"
	"go-echo-gorm-rest/model"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	mockDB, sqlMock, c, rec, _ := Setup(t)

	h := handler.CreateUserHandler(mockDB)

	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	query := "SELECT * FROM `users` WHERE `users`.`id` = ?"

	sqlMock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("1").WillReturnRows(sqlmock.NewRows(model.UserColumns).AddRow(mock.User.ID, mock.User.Name, mock.User.CreatedAt, mock.User.UpdatedAt))

	h.GetUser(c)

	assert.Equal(t, string(mock.UserJson), rec.Body.String())
}
