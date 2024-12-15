package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"github.com/imperatorofdwelling/payment-svc/internal/storage/postgres"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"go.uber.org/zap"
)

type ILogsSvc interface {
	InsertLog(ctx context.Context, payment *yoopayment.Payment) error
}

type LogsSvc struct {
	repo postgres.ILogsRepo
	log  *zap.SugaredLogger
}

func NewLogsService(repo postgres.ILogsRepo, log *zap.SugaredLogger) *LogsSvc {
	return &LogsSvc{repo, log}
}

func (s *LogsSvc) InsertLog(ctx context.Context, payment *yoopayment.Payment) error {
	const op = "service.logs.InsertLog"

	transactionID, err := uuid.Parse(payment.ID)
	if err != nil {
		s.log.Errorf("%s: %v", op, err)
		return fmt.Errorf("%s: %w", op, err)
	}

	//var paymentMethod model.PaymentMethod

	//if pm, ok := payment.PaymentMethod.(model.PaymentMethod); ok {
	//	paymentMethod = pm
	//} else {
	//	s.log.Errorf("%s: %v", op, err)
	//	return fmt.Errorf("%s: %w", op, err)
	//}

	newLog := model.LogPaymentRequest{
		TransactionID: transactionID,
		Status:        string(payment.Status),
		Value:         payment.Amount.Value,
		Currency:      payment.Amount.Currency,
		Type:          "bank_card",
	}

	err = s.repo.InsertLog(ctx, &newLog)
	if err != nil {
		return err
	}

	return nil
}
