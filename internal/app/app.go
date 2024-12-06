package app

import (
	"github.com/imperatorofdwelling/payment-svc/internal/config"
	"github.com/imperatorofdwelling/payment-svc/internal/handler/http"
	"github.com/imperatorofdwelling/payment-svc/pkg/logger"
)

type App struct {
	Server *http.Server
}

func NewApp() *App {
	cfg := config.MustLoad()

	log := logger.NewZapLogger(cfg.Env)

	router := http.NewRouter()

	server := http.NewServer(cfg.Server, router, log)

	app := &App{
		Server: server,
	}

	return app
}
