package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// NewValidator 함수는 model필드를 생성할 때마다 new validator를 생성합니다.
func NewValidator() *validator.Validate {
	//Book model을 위한 validator생성
	validate := validator.New()

	//  uuid.UUID fields를 위한 커스텀 validation체크
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	return validate
}

// ValidatorErrors func으로 validation error를 각 invalid한 field들을 보여주도록 함
func ValidatorErrors(err error) map[string]string {
	// fields맵을 정의
	fields := map[string]string{}

	// 각 invalid fied에 error메세지 생성
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields
}
