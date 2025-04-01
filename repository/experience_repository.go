package repository

import (
	"chat_upgrade/model"

	"gorm.io/gorm"
)

type IExperienceRepository interface {
	CreateExperience(exp *model.Experience) error
	GetAllExperiences() ([]model.Experience, error)
	DeleteExperience(id uint) error
}

type experienceRepository struct {
	db *gorm.DB
}

func NewExperienceRepository(db *gorm.DB) IExperienceRepository {
	return &experienceRepository{db}
}

func (er *experienceRepository) CreateExperience(exp *model.Experience) error {
	return er.db.Create(exp).Error
}

func (er *experienceRepository) GetAllExperiences() ([]model.Experience, error) {
	var experiences []model.Experience
	err := er.db.Find(&experiences).Error
	return experiences, err
}

func (er *experienceRepository) DeleteExperience(id uint) error {
	return er.db.Delete(&model.Experience{}, id).Error
}
