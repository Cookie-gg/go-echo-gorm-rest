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

func TestGetProfile(t *testing.T) {
	s := service.CreateProfileService(mockDB)

	query := "SELECT * FROM `profiles` WHERE `profiles`.`user_id` = ? LIMIT 1"

	sqlMock.MatchExpectationsInOrder(false)
	sqlMock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(sqlmock.NewRows(mock.ProfileColumns).AddRow(mock.Profile.Age, mock.Profile.Facebook, mock.Profile.Gender, mock.Profile.ID, mock.Profile.Twitter, mock.Profile.UserID))

	profile := model.Profile{UserID: 1}

	if assert.NoError(t, s.Get(&profile)) {
		assert.Equal(t, mock.Profile, profile)
	}
}
