package main

import (
	"github.com/imperatorofdwelling/payment-svc/internal/app"
)

func main() {
	appV1 := app.NewApp()

	if err := appV1.Server.Start(); err != nil {
		panic(err)
	}

	appV1.Server.Stop(appV1.Scheduler)
}
