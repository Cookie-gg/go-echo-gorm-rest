package mock

import (
	"encoding/json"
	"fmt"
	"go-echo-gorm-rest/model"
	"time"

	"gorm.io/gorm"
)

var User = model.User{
	ID:        1,
	Name:      "テスト田中",
	CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	DeletedAt: gorm.DeletedAt{Valid: false},
}

var jsonByte, _ = json.Marshal(User)

var UserJson = fmt.Sprintf("%s\n", string(jsonByte))
