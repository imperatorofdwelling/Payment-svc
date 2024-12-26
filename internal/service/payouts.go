package service

import (
	"context"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"github.com/imperatorofdwelling/payment-svc/internal/storage/postgres"
	"go.uber.org/zap"
)

type IPayoutsSvc interface {
	CreatePayout(ctx context.Context, payout model.Payout) error
}

type PayoutsSvc struct {
	repo             postgres.IPayoutsRepo
	logsSvc          ILogsSvc
	payoutSubscriber IPayoutSubscriber
	log              *zap.SugaredLogger
}

func NewPayoutsService(repo postgres.IPayoutsRepo, payoutSubscriber IPayoutSubscriber, logsSvc ILogsSvc, log *zap.SugaredLogger) *PayoutsSvc {
	return &PayoutsSvc{repo, logsSvc, payoutSubscriber, log}
}

func (s *PayoutsSvc) CreatePayout(ctx context.Context, payout model.Payout) error {
	const op = "service.payout.NewPayout"

	newLog := &model.Log{
		TransactionID:   payout.ID,
		TransactionType: model.PayoutType,
		Status:          *payout.Status,
		Value:           payout.Amount.Value,
		Currency:        payout.Amount.Currency,
	}

	err := s.logsSvc.InsertLog(ctx, newLog)
	if err != nil {
		return err
	}

	err = s.payoutSubscriber.Subscribe(payout.ID, *payout.Status)
	if err != nil {
		return err
	}

	return nil
}
