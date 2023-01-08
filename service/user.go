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

func (m *UserService) Create() (user model.User, err error) {
	m.db.Create(&user)
	return
}
func (m *UserService) Get(id string) (user model.User, err error) {
	m.db.First(&user, id)
	return
}
func (m *UserService) Update() (user model.User, err error) {
	m.db.Save(&user)
	return
}
func (m *UserService) Delete() (user model.User, err error) {
	m.db.Delete(&user)
	return
}
