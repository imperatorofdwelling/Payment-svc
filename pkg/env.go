package pkg

import (
	"fmt"
	"strings"
)

type Env string

// String implemented method for flag Var
func (e *Env) String() string {
	return string(*e)
}

// Set implemented method for flag Var
func (e *Env) Set(s string) error {
	upperValue := Env(strings.ToLower(s))

	validEnvironments := []Env{LocalEnv, DevEnv, ProdEnv}

	for _, env := range validEnvironments {
		if env == upperValue {
			*e = upperValue
			return nil
		}
	}
	return fmt.Errorf("invalid environment: %s, valid values: %v", s, validEnvironments)
}

const (
	LocalEnv Env = "local"
	DevEnv   Env = "dev"
	ProdEnv  Env = "prod"
)
