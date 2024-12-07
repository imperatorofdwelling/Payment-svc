package service

import (
	"github.com/imperatorofdwelling/payment-svc/internal/storage/postgres/repo"
	"go.uber.org/zap"
)

type IPaymentSvc interface {
	GetSTest() string
}

type PaymentSvc struct {
	repo repo.IPaymentRepo
	log  *zap.SugaredLogger
}

func NewPaymentSvc(repo repo.IPaymentRepo, log *zap.SugaredLogger) *PaymentSvc {
	return &PaymentSvc{
		repo: repo,
		log:  log,
	}
}

func (s *PaymentSvc) GetSTest() string {
	return s.repo.GetTest()
}
