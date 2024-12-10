package service

import (
	"github.com/imperatorofdwelling/payment-svc/internal/storage/postgres"
	"go.uber.org/zap"
)

type IPaymentSvc interface {
	GetSTest() string
}

type PaymentSvc struct {
	repo postgres.IPaymentRepo
	log  *zap.SugaredLogger
}

func NewPaymentSvc(repo postgres.IPaymentRepo, log *zap.SugaredLogger) *PaymentSvc {
	return &PaymentSvc{
		repo: repo,
		log:  log,
	}
}

func (s *PaymentSvc) GetSTest() string {
	return s.repo.GetTest()
}
