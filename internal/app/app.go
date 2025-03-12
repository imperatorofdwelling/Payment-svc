package app

import (
	"github.com/imperatorofdwelling/payment-svc/internal/config"
	"github.com/imperatorofdwelling/payment-svc/internal/handler/http"
	"github.com/imperatorofdwelling/payment-svc/internal/lib/scheduler"
	v10 "github.com/imperatorofdwelling/payment-svc/internal/lib/validator"
	"github.com/imperatorofdwelling/payment-svc/internal/storage"
	"github.com/imperatorofdwelling/payment-svc/pkg/logger"
)

type App struct {
	Server *http.Server

	Scheduler *scheduler.Scheduler
}

func NewApp() *App {
	cfg := config.MustLoad()

	log := logger.NewZapLogger(cfg.Env)

	s := scheduler.NewScheduler(log)
	s.Run()

	v10.NewValidator(log)

	storages := storage.GetStorages(cfg, log)

	router := http.NewRouter(storages, log, cfg)

	server := http.NewServer(cfg.Server, router.Handler, log)

	app := &App{
		Server:    server,
		Scheduler: s,
	}

	return app
}
