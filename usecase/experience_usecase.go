package usecase

import (
	"chat_upgrade/model"
	"chat_upgrade/repository"
	"mime/multipart"
)

type IExperienceUsecase interface {
	CreateExperience(userID uint, title, techStack, content string, file *multipart.FileHeader) (model.Experience, error)
	GetAllExperiences() ([]model.Experience, error)
	DeleteExperience(id uint) error
}

type experienceUsecase struct {
	er  repository.IExperienceRepository
	s3  repository.IS3Repository
}

func NewExperienceUsecase(er repository.IExperienceRepository, s3 repository.IS3Repository) IExperienceUsecase {
	return &experienceUsecase{er, s3}
}

func (eu *experienceUsecase) CreateExperience(userID uint, title, techStack, content string, file *multipart.FileHeader) (model.Experience, error) {
	// アイコンをS3にアップロード
	iconURL, err := eu.s3.UploadFile(file)
	if err != nil {
		return model.Experience{}, err
	}

	// 経験データを作成
	experience := model.Experience{
		UserID:    userID,
		Title:     title,
		TechStack: techStack,
		Icon:      iconURL,
		Content:   content,
	}

	// データベースに保存
	if err := eu.er.CreateExperience(&experience); err != nil {
		return model.Experience{}, err
	}

	return experience, nil
}

func (eu *experienceUsecase) GetAllExperiences() ([]model.Experience, error) {
	return eu.er.GetAllExperiences()
}

func (eu *experienceUsecase) DeleteExperience(id uint) error {
	return eu.er.DeleteExperience(id)
}
