package logger

import (
	"github.com/imperatorofdwelling/payment-svc/pkg"
	"go.uber.org/zap"
)

func NewZapLogger(env pkg.Env) *zap.SugaredLogger {
	var logger *zap.Logger

	switch env {
	case pkg.ProdEnv:
		logger, _ = zap.NewProduction()
	default:
		logger, _ = zap.NewDevelopment()
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	if env != pkg.ProdEnv {
		sugar.Infof("Dev mode logger enabled with env: %s", env)
	}

	return sugar
}
