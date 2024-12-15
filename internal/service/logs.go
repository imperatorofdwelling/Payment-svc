package service

import (
	"github.com/imperatorofdwelling/payment-svc/internal/storage/postgres"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"go.uber.org/zap"
)

type ILogsSvc interface {
	InsertLog(payment *yoopayment.Payment)
}

type LogsSvc struct {
	repo postgres.ILogsRepo
	log  *zap.SugaredLogger
}

func NewLogsService(repo postgres.ILogsRepo, log *zap.SugaredLogger) *LogsSvc {
	return &LogsSvc{repo, log}
}

func (l *LogsSvc) InsertLog(payment *yoopayment.Payment) {
	const op = "service.logs.InsertLog"
}
