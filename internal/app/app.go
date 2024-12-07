package app

import (
	"github.com/imperatorofdwelling/payment-svc/internal/config"
	"github.com/imperatorofdwelling/payment-svc/internal/handler/http"
	"github.com/imperatorofdwelling/payment-svc/internal/storage"
	"github.com/imperatorofdwelling/payment-svc/pkg/logger"
)

type App struct {
	Server *http.Server
}

func NewApp() *App {
	cfg := config.MustLoad()

	log := logger.NewZapLogger(cfg.Env)

	storages := storage.GetStorages(cfg, log)

	router := http.NewRouter(storages, log)

	server := http.NewServer(cfg.Server, router.Handler, log)

	app := &App{
		Server: server,
	}

	return app
}
