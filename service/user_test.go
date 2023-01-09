package service_test

import (
	"go-echo-gorm-rest/mock"
	"go-echo-gorm-rest/model"
	"go-echo-gorm-rest/service"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
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

func TestCreateUser(t *testing.T) {
	s := service.CreateUserService(mockDB)

	query := "INSERT INTO `users`"

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(regexp.QuoteMeta(query)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	assert.NoError(t, s.Create(&mock.User))
}

func TestGetUser(t *testing.T) {
	s := service.CreateUserService(mockDB)

	query := "SELECT * FROM `users` WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL LIMIT 1"

	sqlMock.MatchExpectationsInOrder(false)
	sqlMock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(sqlmock.NewRows(model.UserColumns).AddRow(mock.User.ID, mock.User.Name, mock.User.CreatedAt, mock.User.UpdatedAt, mock.User.DeletedAt))

	user := model.User{}

	if assert.NoError(t, s.Get(&user, 1)) {
		assert.Equal(t, mock.User, user)
	}
}
