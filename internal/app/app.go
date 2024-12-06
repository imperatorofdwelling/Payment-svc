package app

import (
	"github.com/imperatorofdwelling/payment-svc/internal/config"
	"github.com/imperatorofdwelling/payment-svc/internal/handler/http"
)

type App struct {
	Server *http.Server
}

func NewApp() *App {
	cfg := config.MustLoad()

	router := http.NewRouter()

	server := http.NewServer(cfg.Server, router)

	app := &App{
		Server: server,
	}

	return app
}
