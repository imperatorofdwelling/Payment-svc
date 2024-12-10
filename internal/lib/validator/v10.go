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
	registerValidation("required_if", requiredIfAnotherFieldContainsValidator, log, newValidate)
	//registerValidation("e164", e164Validator, log, newValidate)

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

//func e164Validator(log *zap.SugaredLogger) validator.Func {
//	return func(fl validator.FieldLevel) bool {
//		v := fl.Field()
//
//		if v.Kind() != reflect.String {
//			log.Errorf("value %v is not a string", v)
//			return false
//		}
//
//		e164Regex := `^\+[1-9]\d{1,14}$`
//
//		reg := regexp.MustCompile(e164Regex)
//		if !reg.MatchString(v.String()) {
//			log.Errorf("value %v is not a valid e164", v)
//			return false
//		}
//
//		return true
//	}
//}

func requiredIfAnotherFieldContainsValidator(log *zap.SugaredLogger) validator.Func {
	return func(fl validator.FieldLevel) bool {
		v := fl.Field()

		if v.Kind() != reflect.String {
			log.Errorf("value %v is not a string", v)
			return false
		}

		param := fl.Param()

		paramSplit := strings.Split(param, ":")
		if len(paramSplit) != 2 {
			log.Errorf("value %v has incorect format", v)
			return false
		}

		key := paramSplit[0]
		value := paramSplit[1]

		comparedFiledValue := fl.Parent().FieldByName(key).String()
		if comparedFiledValue == "" {
			log.Errorf("value %v has not compared", v)
			return false
		}

		if strings.Contains(value, "|") {
			splitValue := strings.Split(value, "|")

			isEqual := false

			for _, val := range splitValue {
				if val == comparedFiledValue {
					isEqual = true
				}
			}

			if !isEqual {
				log.Errorf("value %v is not a valid field", v)
				return false
			}

			return true
		}

		if comparedFiledValue != value && v.String() != "" {
			log.Errorf("value %v is not a valid field, it should be %s", v, value)
			return false
		}

		return true
	}
}
