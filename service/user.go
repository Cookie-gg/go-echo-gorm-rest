package service

import (
	"go-echo-gorm-rest/model"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func CreateUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}

func (m *UserService) Create(user *model.User) error {
	return m.db.Debug().Create(user).Error
}
func (m *UserService) Get(user *model.User, filters ...interface{}) error {
	return m.db.Debug().Take(user, filters).Error
}
func (m *UserService) Update(user *model.User) error {
	return m.db.Save(&user).Error
}
func (m *UserService) Delete(user *model.User) error {
	return m.db.Delete(&user).Error
}
