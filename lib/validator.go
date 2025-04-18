package lib

import "github.com/go-playground/validator/v10"

type Validator struct {
	validate *validator.Validate
}

func (v Validator) Validate(obj any) error {
	return v.validate.Struct(obj)
}

func NewValidator() Validator {
	return Validator{validate: validator.New(validator.WithRequiredStructEnabled())}
}
