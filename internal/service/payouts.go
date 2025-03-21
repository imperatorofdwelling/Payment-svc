package service

import (
	"context"
	"fmt"
	"github.com/eclipsemode/go-yookassa-sdk/yookassa/model"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"github.com/imperatorofdwelling/payment-svc/internal/lib/scheduler"
	"github.com/imperatorofdwelling/payment-svc/internal/storage/postgres"
	"go.uber.org/zap"
)

type IPayoutsSvc interface {
	CreatePayout(ctx context.Context, payout yoomodel.Payout) error
	SchedulePayout(ctx context.Context, payout yoomodel.Payout, scheduler *scheduler.Scheduler) error
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

func (s *PayoutsSvc) CreatePayout(ctx context.Context, payout yoomodel.Payout) error {
	const op = "service.payout.NewPayout"

	newLog := &model.Log{
		TransactionID:   payout.ID,
		TransactionType: yoomodel.PayoutType,
		Status:          payout.Status,
		Value:           payout.Amount.Value,
		Currency:        payout.Amount.Currency,
	}

	err := s.logsSvc.InsertLog(ctx, newLog)
	if err != nil {
		return err
	}

	err = s.payoutSubscriber.Subscribe(payout.ID, payout.Status)
	if err != nil {
		return err
	}

	return nil
}

func (s *PayoutsSvc) SchedulePayout(ctx context.Context, payout yoomodel.Payout, scheduler *scheduler.Scheduler) error {
	const op = "service.payout.SchedulePayout"

	// TODO get payout date

	scheduler.Create("0 22 * * *", func() {
		fmt.Println("New payout")
	})
	return nil
}
