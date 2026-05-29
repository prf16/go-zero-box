package tools

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{validate: validator.New()}
}

func (v *Validator) Validate(r *http.Request, data any) error {
	return v.validate.Struct(data)
}
