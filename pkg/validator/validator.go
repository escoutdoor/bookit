package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unicode"

	"github.com/go-playground/validator/v10"
)

var (
	ErrInvalidDateFormat = errors.New("invalid date format")
)

type Validator struct {
	v    *validator.Validate
	Errm map[string]string
}

func New() *Validator {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	validate.RegisterValidation("password", validatePw)

	return &Validator{
		v:    validate,
		Errm: make(map[string]string),
	}
}

func (vl *Validator) IsValid() bool {
	return len(vl.Errm) == 0
}

func (vl *Validator) AddErr(key, msg string) {
	_, ok := vl.Errm[key]
	if !ok {
		vl.Errm[key] = msg
	}
}

func (vl *Validator) Validate(b interface{}) {
	err := vl.v.Struct(b)
	if err != nil {
		for _, item := range err.(validator.ValidationErrors) {
			ve := vl.getValidationErr(item)
			vl.Errm[item.Field()] = ve.Error()
		}
	}
}

func (vl *Validator) ParseDate(dobStr string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", dobStr)
	if err != nil {
		return time.Time{}, ErrInvalidDateFormat
	}
	return t, nil
}

func (vl *Validator) getValidationErr(err validator.FieldError) error {
	var (
		field = err.Field()
		param = err.Param()
	)

	switch err.Tag() {
	case "required":
		return fmt.Errorf("field %s is required", field)
	case "min":
		return fmt.Errorf("field %s must be at least %s characters long", field, param)
	case "max":
		return fmt.Errorf("field %s must be at most %s characters long", field, param)
	case "email":
		return fmt.Errorf("field %s must be a valid email address", field)
	case "url":
		return fmt.Errorf("field %s must be a valid URL", field)
	case "uuid":
		return fmt.Errorf("field %s must be a valid UUID", field)
	case "gte":
		return fmt.Errorf("field %s must be greater than or equal to %s", field, param)
	case "gt":
		return fmt.Errorf("field %s must be greater than %s", field, param)
	case "lte":
		return fmt.Errorf("field %s must be less than or equal to %s", field, param)
	case "e164":
		return fmt.Errorf("field %s must be a valid phone number", field)
	default:
		return fmt.Errorf("field %s is invalid", field)
	}
}

func (vl *Validator) fieldName(s string) string {
	parts := strings.Split(s, ".")
	if len(s) > 1 {
		return strings.ToLower(strings.Join(parts[1:], "."))
	}

	return strings.ToLower(s)
}

func validatePw(fl validator.FieldLevel) bool {
	var (
		pw         = fl.Field().String()
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for _, char := range pw {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSpecial
}
