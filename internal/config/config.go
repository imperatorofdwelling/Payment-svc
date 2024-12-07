package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/imperatorofdwelling/payment-svc/pkg"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Env    pkg.Env `yaml:"env" env-default:"local" env-required:"true"`
	Server `yaml:"server" env-required:"true"`
	Db
}

type Server struct {
	Host        string        `yaml:"host" env-default:"localhost" env-required:"true"`
	Port        int           `yaml:"port" env-default:"8080" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Db struct {
	Postgres `yaml:"postgres" env-required:"true"`
	Redis    `yaml:"redis" env-required:"true"`
}

type Postgres struct {
	Host     string `yaml:"host" env-default:"localhost" env-required:"true"`
	Username string `yaml:"username" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	DbName   string `yaml:"db_name" env-required:"true"`
	Port     int    `yaml:"port" env-default:"5432" env-required:"true"`
	SSLMode  string `yaml:"ssl_mode"`
}

type Redis struct {
	Password string `yaml:"password" env-required:"true"`
}

type flagsOption struct {
	Env pkg.Env
}

func MustLoad() *Config {
	flags, err := loadFlags()
	if err != nil {
		panic(err)
	}

	cfg, err := loadConfig(flags.Env)
	if err != nil {
		panic(err)
	}

	return cfg
}

func loadConfig(env pkg.Env) (*Config, error) {
	configPath, err := filepath.Abs(fmt.Sprintf("./config/%s.conf.yml", env))
	if err != nil {
		return nil, fmt.Errorf("can't get absolute path for config file: %w", err)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file %s does not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("can't read config file %s: %w", configPath, err)
	}

	return &cfg, nil
}

func loadFlags() (flagsOption, error) {
	var currEnv pkg.Env
	flag.Var(&currEnv, "env", "Environment type")
	flag.Parse()

	fOption := flagsOption{
		Env: currEnv,
	}

	return fOption, nil
}
