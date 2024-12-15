package service

import (
	"github.com/imperatorofdwelling/payment-svc/internal/storage/postgres"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"go.uber.org/zap"
)

type IPaymentSvc interface {
	CreatePayment(*yoopayment.Payment) *yoopayment.Payment
	SavePaymentLog(*yoopayment.Payment)
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

func (s *PaymentSvc) CreatePayment(*yoopayment.Payment) *yoopayment.Payment {
	return nil
}
