package validator

import (
	"go-rest-api/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserValidate(user models.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user models.User) error {
	return validation.ValidateStruct(
		&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 30).Error("limited max 30 length"),
			is.Email.Error("invalid email format"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 30).Error("password length must be greater than 6 and less than 30"),
		),
	)
}
