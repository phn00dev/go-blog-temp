package errors

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/phn00dev/go-blog-temp/internal/providers/validation"
	"strings"
)

var errorsList = make(map[string]string)

func Init() {
	errorsList = map[string]string{}
}

func SetFormErrors(err error) {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, failedErr := range validationErrors {
			Add(failedErr.Field(), GetFormMsg(failedErr.Tag()))
		}
	}
}

func Add(key, value string) {
	errorsList[strings.ToLower(key)] = value
}

func GetFormMsg(tag string) string {
	return validation.ErrorMessages()[tag]
}

func Get() map[string]string {
	return errorsList
}
