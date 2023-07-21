package tools

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func Validation_errors(schema any) []string {
	validate := validator.New()
	err := validate.Struct(schema)

	if err != nil {
		var result []string

		for _, e := range err.(validator.ValidationErrors) {
			result = append(result, strings.Split(e.Error(), "Error:")[1])
		}

		return result
	}

	return nil
}
