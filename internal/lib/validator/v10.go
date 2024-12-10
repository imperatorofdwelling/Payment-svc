package v10

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"reflect"
	"regexp"
)

var Validate *validator.Validate

func NewValidator(log *zap.SugaredLogger) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.RegisterValidation("currency", currencyValidator(log))
	if err != nil {
		log.Errorf("failed to register currency validation: %s", err.Error())
	}

	Validate = validate
}

func currencyValidator(log *zap.SugaredLogger) validator.Func {
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
