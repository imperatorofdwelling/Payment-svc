package service

import (
	"context"
	"fmt"
	"github.com/eclipsemode/go-yookassa-sdk/yookassa/model"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"github.com/imperatorofdwelling/payment-svc/internal/storage/postgres"
	"go.uber.org/zap"
)

type IPaymentSvc interface {
	CreatePayment(context.Context, *yoomodel.Payment) error
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

func (s *PaymentSvc) CreatePayment(ctx context.Context, payment *yoomodel.Payment) error {
	const op = "service.payments.CreatePayment"

	newLog := &model.Log{
		TransactionID:   payment.ID,
		TransactionType: yoomodel.PaymentType,
		Status:          payment.Status,
		Value:           payment.Amount.Value,
		Currency:        payment.Amount.Currency,
	}

	err := s.logsSvc.InsertLog(ctx, newLog)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
