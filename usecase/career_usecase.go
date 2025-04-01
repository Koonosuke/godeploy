package usecase

import (
	"chat_upgrade/model"
	"chat_upgrade/repository"
)

type ICareerUsecase interface {
	CreateCareer(career *model.Career) error
	GetCareers() ([]model.Career, error)
	DeleteCareer(id int) error
}

type CareerUsecase struct {
	repo repository.ICareerRepository
}

func NewCareerUsecase(repo repository.ICareerRepository) *CareerUsecase {
	return &CareerUsecase{repo: repo}
}

func (u *CareerUsecase) CreateCareer(career *model.Career) error {
	return u.repo.Create(career)
}

func (u *CareerUsecase) GetCareers() ([]model.Career, error) {
	return u.repo.GetAll()
}

func (u *CareerUsecase) DeleteCareer(id int) error {
	return u.repo.Delete(id)
}
