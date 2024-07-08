package validator

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var vd *validator.Validate

// regex
var (
	yearRegex = regexp.MustCompile(`^(\d{4})$`)
)

func ValidateYear(s int) bool {
	return yearRegex.Match([]byte(fmt.Sprintf("%d", s)))
}

func InitValidator() {
	vd = validator.New()
	// custom validator
	vd.RegisterValidation("year", func(fl validator.FieldLevel) bool {
		return ValidateYear(int(fl.Field().Int()))
	})
}

func StructValidator(arg interface{}) []string {
	// struct validator
	err := vd.Struct(arg)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); !ok {
			fields := make([]string, 0)
			for _, err := range err.(validator.ValidationErrors) {
				fieldName := strings.ToLower(string(err.Field()[0])) + err.Field()[1:]
				errMsg := fmt.Sprintf("'%s' failed on the rule: '%s'", fieldName, err.ActualTag())
				switch err.Tag() {
				case "year":
					errMsg = fmt.Sprintf("'%s' must have 4 digits", fieldName)
				}
				fields = append(fields, errMsg)
			}
			return fields
		}
	}
	return nil
}
