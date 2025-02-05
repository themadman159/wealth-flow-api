package validateutil

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Validate(req interface{}) error {
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			errorMessages = append(errorMessages, fmt.Sprintf("Field %s is required", fieldName))
		}

		return fmt.Errorf(strings.Join(errorMessages, ", "))
	}
	return nil
}
