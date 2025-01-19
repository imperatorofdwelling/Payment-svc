package app

import (
	"github.com/imperatorofdwelling/payment-svc/internal/config"
	"github.com/imperatorofdwelling/payment-svc/internal/handler/http"
	kafka "github.com/imperatorofdwelling/payment-svc/internal/handler/kafka/consumer"
	v10 "github.com/imperatorofdwelling/payment-svc/internal/lib/validator"
	"github.com/imperatorofdwelling/payment-svc/internal/storage"
	"github.com/imperatorofdwelling/payment-svc/pkg/logger"
)

type App struct {
	Server *http.Server
}

func NewApp() *App {
	cfg := config.MustLoad()

	log := logger.NewZapLogger(cfg.Env)

	v10.NewValidator(log)

	storages := storage.GetStorages(cfg, log)

	router := http.NewRouter(storages, log, cfg)

	kafka.SetupKafkaConsumers()

	server := http.NewServer(cfg.Server, router.Handler, log)

	app := &App{
		Server: server,
	}

	return app
}
