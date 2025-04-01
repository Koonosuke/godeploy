package repository

import (
	"chat_upgrade/model"

	"gorm.io/gorm"
)

type ICareerRepository interface {
	Create(career *model.Career) error
	GetAll() ([]model.Career, error)
	Delete(id int) error
}

type CareerRepository struct {
	db *gorm.DB
}

func NewCareerRepository(db *gorm.DB) *CareerRepository {
	return &CareerRepository{db: db}
}

func (r *CareerRepository) Create(career *model.Career) error {
	return r.db.Create(career).Error
}

func (r *CareerRepository) GetAll() ([]model.Career, error) {
	var careers []model.Career
	err := r.db.Order("created_at DESC").Find(&careers).Error
	return careers, err
}

func (r *CareerRepository) Delete(id int) error {
	return r.db.Delete(&model.Career{}, id).Error
}
