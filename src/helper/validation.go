package helper

import (
	"golang-test/common"

	"github.com/go-playground/validator/v10"
)

func NewValidation(validate *validator.Validate) Validation {
	return &validationImpl{
		validate: validate,
	}
}

type Validation interface {
	ValidationRequestRead(rl *common.WebRequestRead) error
	ValidationRequestCreate(rl *common.WebRequestCreate) error
	ValidationRequestUpdate(rl *common.WebRequestUpdate) error
	ValidationRequestDelete(rl *common.WebRequestDelete) error
}

type validationImpl struct {
	validate *validator.Validate
}

func (v *validationImpl) ValidationRequestRead(rl *common.WebRequestRead) error {
	return v.validate.Struct(rl)
}

func (v *validationImpl) ValidationRequestCreate(rl *common.WebRequestCreate) error {
	return v.validate.Struct(rl)
}

func (v *validationImpl) ValidationRequestUpdate(rl *common.WebRequestUpdate) error {
	return v.validate.Struct(rl)
}

func (v *validationImpl) ValidationRequestDelete(rl *common.WebRequestDelete) error {
	return v.validate.Struct(rl)
}
