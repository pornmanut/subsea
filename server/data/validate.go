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

//Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

// Validation is a CustomValidator
type Validation struct {
	validate *validator.Validate
}

// NewValidation create a new Validation type for validate sturct
func NewValidation() *Validation {
	validator := validator.New()
	return &Validation{validate: validator}
}

// Validate the item
// for more detail the returned error can be cast into a
// validator.ValidationErrors collection
//

// Validate is methods for validation given by interface struct return mutil errors
func (v *Validation) Validate(i interface{}) ValidationErrors {

	err := v.validate.Struct(i)
	if ves, ok := err.(validator.ValidationErrors); ok {
		if len(ves) == 0 {
			return nil
		}

		var returnErrs []ValidationError

		for _, ve := range ves {
			// cast the FieldError into ValidationError append into slice
			// from implemented Error()
			e := ValidationError{ve.(validator.FieldError)}
			returnErrs = append(returnErrs, e)
		}
		return returnErrs
	}
	return nil
}
