package validate

import "github.com/go-playground/validator/v10"

var (
	validate = validator.New()
)

// Validate validates the input
func Validate(i interface{}) error {
	return validate.Struct(i)
}
