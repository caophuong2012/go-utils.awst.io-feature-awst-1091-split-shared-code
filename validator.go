package utils

import "github.com/go-playground/validator/v10"

func newValidator() *validator.Validate {
	return validator.New()
}

func Validate(myStruct interface{}) error {
	return newValidator().Struct(myStruct)
}
