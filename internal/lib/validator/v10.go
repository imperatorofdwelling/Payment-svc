package v10

import (
	"github.com/go-playground/validator/v10"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"go.uber.org/zap"
	"reflect"
	"regexp"
	"strings"
)

var Validate *validator.Validate

func NewValidator(log *zap.SugaredLogger) {
	newValidate := validator.New(validator.WithRequiredStructEnabled())

	registerValidation("money", moneyValidator, log, newValidate)
	registerValidation("currency", currencyValidator, log, newValidate)
	registerValidation("should_exist_field", shouldExistFieldsValidator, log, newValidate)

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

func currencyValidator(log *zap.SugaredLogger) validator.Func {
	return func(fl validator.FieldLevel) bool {
		v := fl.Field()

		if v.Kind() != reflect.String {
			log.Errorf("value %v is not a string", v)
			return false
		}

		if _, exists := model.ValidCurrencies[v.String()]; !exists {
			log.Errorf("value %v is not a valid currency", v)
			return false
		}

		return true
	}
}

func shouldExistFieldsValidator(log *zap.SugaredLogger) validator.Func {
	return func(fl validator.FieldLevel) bool {
		v := fl.Field()

		if v.Kind() != reflect.String {
			log.Errorf("value %v is not a string", v)
			return false
		}

		param := fl.Param()
		words := strings.Fields(param)

		result := make(map[string]string)

		for i := 0; i < len(words)-1; i += 2 {
			key := words[i]
			value := words[i+1]
			result[key] = value
		}

		parent := fl.Parent()

		for fieldKey, fieldNameValue := range result {
			shouldExistValue := parent.FieldByName(fieldNameValue)

			if fieldKey != v.String() {
				continue
			}
			//log.Debugf("AAAAAAAAAAAAAAAAAAAA %s", fieldKey)

			if shouldExistValue.IsZero() || shouldExistValue.String() == "" {
				return false
			}
		}

		return true
	}
}
