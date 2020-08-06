package data

import (
	"fmt"

	"github.com/go-playground/validator"
)

// ValidationError warps the validators FieldError
type ValidationError struct {
	validator.FieldError
}

// Implement Error
func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' on '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Validation is a CustomValidator
type Validation struct {
	validate *validator.Validate
}

// NewValidation create a new Validation type
func NewValidation() Validation {
	validator := validator.New()

	// TODO
	// register more
	// - email dulplicate
	return Validation{validate: validator}
}

// Validate is methods for validation given by interface struct return mutil errors
func (v *Validation) Validate(i interface{}) ValidationErrors {

	errs := v.validate.Struct(i).(validator.ValidationErrors)

	// not found error
	if len(errs) == 0 {
		return nil
	}

	var returnErrs []ValidationError

	for _, err := range errs {
		// cast the FieldError into ValidationError append into slice

		// type assertion
		// err is error that contain interface that has method Error()
		// we assgin to ValidationError. that has the same method Error()
		// that we implemnt on above

		ve := ValidationError{err.(ValidationError)}
		returnErrs = append(returnErrs, ve)
	}
	return returnErrs
}
