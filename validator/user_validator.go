package validator

import (
	"chat_upgrade/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IUserValidator interface {
	ValidateUser(user model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) ValidateUser(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Username,
			validation.Required.Error("username is required"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
		),
	)
}
