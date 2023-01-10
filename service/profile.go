package service

import (
	"go-echo-gorm-rest/model"

	"gorm.io/gorm"
)

type ProfileService struct {
	db *gorm.DB
}

func CreateProfileService(db *gorm.DB) *ProfileService {
	return &ProfileService{db}
}

func (m *ProfileService) Get(profile *model.Profile) error {
	return m.db.Debug().Where(profile).Take(profile).Error
}
func (m *ProfileService) Update(profile *model.Profile) error {
	return m.db.Save(&profile).Error
}
