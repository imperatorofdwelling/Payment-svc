package service

import (
	"context"
	"fmt"
	"github.com/imperatorofdwelling/payment-svc/internal/storage/postgres"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"go.uber.org/zap"
)

type IPaymentSvc interface {
	CreatePayment(context.Context, *yoopayment.Payment) error
}

type PaymentSvc struct {
	repo    postgres.IPaymentRepo
	log     *zap.SugaredLogger
	logsSvc ILogsSvc
}

func NewPaymentSvc(repo postgres.IPaymentRepo, logsSvc ILogsSvc, log *zap.SugaredLogger) *PaymentSvc {
	return &PaymentSvc{
		repo,
		log,
		logsSvc,
	}
}

func (s *PaymentSvc) CreatePayment(ctx context.Context, payment *yoopayment.Payment) error {
	const op = "service.payments.CreatePayment"

	err := s.logsSvc.InsertLog(ctx, payment)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
