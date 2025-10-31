package errors

import (
	"blog/internal/providers/validation"
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

var errorList = make(map[string]string)

func Init() {
	errorList = map[string]string{}
}

func SetErrorList(err error) {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, err := range validationErrors {
			errorList[strings.ToLower(err.Field())] = GetErrorMsg(err.Tag())
		}
	}

}

func GetErrorMsg(tag string) string {
	return validation.ErrorMessage()[tag]
}

func GetErrorList() map[string]string {
	return errorList
}
