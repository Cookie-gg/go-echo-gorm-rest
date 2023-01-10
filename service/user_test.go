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

func TestCreateUser(t *testing.T) {
	s := service.CreateUserService(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users`")).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `profiles`")).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	assert.NoError(t, s.Create(&mock.UserWithProfile))
}

func TestGetUser(t *testing.T) {
	s := service.CreateUserService(mockDB)

	query := "SELECT * FROM `profiles` WHERE `profiles`.`user_id` = ?"

	sqlMock.MatchExpectationsInOrder(false)
	sqlMock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(sqlmock.NewRows(mock.ProfileColumns).AddRow(mock.Profile.Age, mock.Profile.Facebook, mock.Profile.Gender, mock.Profile.ID, mock.Profile.Twitter, mock.Profile.UserID))

	query = "SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = ? LIMIT 1"

	sqlMock.MatchExpectationsInOrder(false)
	sqlMock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(sqlmock.NewRows(mock.UserColumns).AddRow(mock.User.CreatedAt, mock.User.DeletedAt, mock.User.ID, mock.User.Name, mock.User.UpdatedAt))

	user := model.User{ID: 1}

	if assert.NoError(t, s.Get(&user)) {
		assert.Equal(t, mock.User, user)
	}
}
