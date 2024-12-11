package v10

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"reflect"
	"regexp"
	"strings"
)

var Validate *validator.Validate

func NewValidator(log *zap.SugaredLogger) {
	newValidate := validator.New(validator.WithRequiredStructEnabled())

	registerValidation("money", moneyValidator, log, newValidate)
	registerValidation("omit_with", omitOptionValidator, log, newValidate)

	Validate = newValidate
}

func registerValidation(tag string, fn func(*zap.SugaredLogger) validator.Func, log *zap.SugaredLogger, validate *validator.Validate) {
	err := validate.RegisterValidation(tag, fn(log))
	if err != nil {
		log.Errorf("failed to register validation '%s': %s", tag, err.Error())
	}
}

func moneyValidator(log *zap.SugaredLogger) validator.Func {
	return func(fl validator.FieldLevel) bool {
		v := fl.Field()

		if v.Kind() != reflect.String {
			log.Errorf("value %v is not a string", v)
			return false
		}

		pattern := `^\d+(\.\d{2})?$`

		reg := regexp.MustCompile(pattern)

		if !reg.MatchString(v.String()) {
			log.Errorf("value %v is not a valid currency", v)
			return false
		}

		return true
	}
}

// omitOptionValidator describes field is no required, but in can only used with chosen field values in struct.
func omitOptionValidator(log *zap.SugaredLogger) validator.Func {
	return func(fl validator.FieldLevel) bool {
		v := fl.Field()

		if v.Kind() != reflect.String {
			log.Errorf("value %v is not a string", v)
			return false
		}

		param := fl.Param()

		arrParam := strings.Split(param, " ")
		key := arrParam[0]
		value := arrParam[1]

		parent := fl.Parent()

		if parent.Kind() != reflect.Struct {
			log.Errorf("value %v is not a struct", v)
			return false
		}

		if v.IsZero() && v.String() != parent.FieldByName(key).String() {
			return true
		}

		if parent.FieldByName(key).String() != value {
			return false
		}

		return true
	}
}
