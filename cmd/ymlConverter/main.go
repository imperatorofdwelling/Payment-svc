package main

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/imperatorofdwelling/payment-svc/internal/config"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type env string

func (e *env) String() string {
	return string(*e)
}

func (e *env) Set(s string) error {
	upperValue := env(strings.ToLower(s))

	validEnvironments := []env{localEnv, devEnv, prodEnv}

	for _, env := range validEnvironments {
		if env == upperValue {
			*e = upperValue
			return nil
		}
	}
	return fmt.Errorf("invalid environment: %s, valid values: %v", s, validEnvironments)
}

const (
	localEnv env = "local"
	devEnv   env = "dev"
	prodEnv  env = "prod"
)

func main() {
	currEnv := loadEnvFlag()

	configPath, err := filepath.Abs(fmt.Sprintf("./config/%s.conf.yml", currEnv))
	if err != nil {
		panic("can't get absolute path for config file")
	}

	var cfg config.Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("can't read config file")
	}

	envFile := ".env"

	file, err := os.OpenFile(envFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	if err := writeConfigToEnvFile(file, &cfg, ""); err != nil {
		fmt.Println("Ошибка при записи в .env файл:", err)
		return
	}

	fmt.Println(".env файл успешно создан и данные записаны.")
}

func loadEnvFlag() env {
	var envFlag env
	flag.Var(&envFlag, "env", "Environment type")
	flag.Parse()

	return envFlag
}

func writeConfigToEnvFile(file *os.File, cfg interface{}, parentName string) error {
	val := reflect.ValueOf(cfg).Elem()
	typ := reflect.TypeOf(cfg).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		fieldName := fieldType.Tag.Get("yaml")
		if fieldName == "" {
			fieldName = fieldType.Name
		}

		if parentName != "" {
			fieldName = parentName + "_" + fieldName
		}

		if field.Kind() == reflect.Struct {
			if err := writeConfigToEnvFile(file, field.Addr().Interface(), fieldType.Name); err != nil {
				return err
			}
		} else {
			line := fmt.Sprintf("%s=%v\n", strings.ToUpper(fieldName), field.Interface())
			if _, err := file.WriteString(line); err != nil {
				return err
			}
		}
	}
	return nil
}
