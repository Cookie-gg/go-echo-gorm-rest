package service_test

import (
	"go-echo-gorm-rest/mock"
	"go-echo-gorm-rest/model"
	"go-echo-gorm-rest/service"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	mockDB, sqlMock := Setup(t)

	s := service.CreateUserService(mockDB)

	query := "SELECT * FROM `users` WHERE `users`.`id` = ?"

	sqlMock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("1").WillReturnRows(sqlmock.NewRows(model.UserColumns).AddRow(mock.User.ID, mock.User.Name, mock.User.CreatedAt, mock.User.UpdatedAt))

	user, _ := s.Get("1")

	assert.Equal(t, mock.User, user)
}
