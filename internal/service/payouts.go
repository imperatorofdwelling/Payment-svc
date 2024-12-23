package service

import (
	"github.com/imperatorofdwelling/payment-svc/internal/storage/postgres"
	"go.uber.org/zap"
)

type IPayoutsSvc interface {
}

type PayoutsSvc struct {
	repo postgres.IPayoutsRepo
	log  *zap.SugaredLogger
}

func NewPayoutsService(repo postgres.IPayoutsRepo, log *zap.SugaredLogger) *PayoutsSvc {
	return &PayoutsSvc{repo, log}
}
