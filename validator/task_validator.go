package validator

import (
	"go-rest-api/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITaskValidator interface {
	TaskValidate(task models.Task) error
}

type taskValidator struct{}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidate(task models.Task) error {
	return validation.ValidateStruct(
		&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 10).Error("limited max 10 length"),
		),
	)
}
