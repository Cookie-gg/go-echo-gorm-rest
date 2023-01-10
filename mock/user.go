package mock

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-echo-gorm-rest/model"
	"go-echo-gorm-rest/util"
	"sort"
	"time"

	"gorm.io/gorm"
)

var User = model.User{
	ID:        1,
	Name:      "テスト田中",
	CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	DeletedAt: gorm.DeletedAt{Valid: false},
	Profile:   Profile,
}

var Profile = model.Profile{
	ID:       1,
	Age:      18,
	Gender:   "male",
	Twitter:  sql.NullString{String: "@test_tanaka", Valid: true},
	Facebook: sql.NullString{String: "テスト田中", Valid: true},
	UserID:   1,
}

var userJsonByte, _ = json.Marshal(User)

var UserJson = fmt.Sprintf("%s\n", string(userJsonByte))

var UserColumns = util.Filter(util.GetJsonKyes(userJsonByte), func(key string) bool {
	return key != "profile"
})

var profileJsonByte, _ = json.Marshal(Profile)

var ProfileJson = fmt.Sprintf("%s\n", string(profileJsonByte))

var ProfileColumns = util.GetJsonKyes(profileJsonByte)

func init() {
	sort.SliceStable(UserColumns, func(i, j int) bool {
		return UserColumns[i] < UserColumns[j]
	})
	sort.SliceStable(ProfileColumns, func(i, j int) bool {
		return ProfileColumns[i] < ProfileColumns[j]
	})
}
