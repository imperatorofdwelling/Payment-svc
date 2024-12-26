package service

import (
	"context"
	"fmt"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"github.com/imperatorofdwelling/payment-svc/internal/storage/redis"
	"github.com/imperatorofdwelling/payment-svc/pkg/json"
	"github.com/imperatorofdwelling/payment-svc/pkg/yookassa"
	"github.com/pkg/errors"
	"math"
	"time"
)

type IPayoutSubscriber interface {
	Subscribe(payoutID string, status model.TransactionStatus) error
}

type PayoutSubscriber struct {
	rdbTransaction     redis.ITransactionRepo
	logsSvc            ILogsSvc
	yookassaPayoutsHdl *yookassa.PayoutsHandler
}

func NewPayoutSubscriber(rdbTransaction redis.ITransactionRepo, logsSvc ILogsSvc, yookassaPayoutsHdl *yookassa.PayoutsHandler) *PayoutSubscriber {
	return &PayoutSubscriber{rdbTransaction, logsSvc, yookassaPayoutsHdl}
}

func (s *PayoutSubscriber) Subscribe(payoutID string, status model.TransactionStatus) error {
	const op = "service.payout.Subscribe"
	if status == model.Succeeded || status == model.Canceled {
		return nil
	}

	isExists, err := s.rdbTransaction.IsExists(payoutID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	if isExists {
		return nil
	}

	err = s.rdbTransaction.Commit(payoutID, status)
	if err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}

	// TODO fix errors in updater
	go s.runUpdater(payoutID)

	return nil
}

func (s *PayoutSubscriber) runUpdater(payoutID string) error {
	const op = "service.payout.runUpdater"

	ch := make(chan struct{}, 1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go signaller(ch, ctx)

	for range ch {
		fmt.Println("TICK")
		var payout model.Payout

		res, err := s.yookassaPayoutsHdl.GetPayoutInfo(payoutID)
		if err != nil {
			return fmt.Errorf("%s: %v", op, err)
		}

		err = json.Read(res.Body, &payout)
		if err != nil {
			return fmt.Errorf("%s: %v", op, err)
		}

		statusInRedis, err := s.rdbTransaction.GetStatus(payoutID)
		if err != nil {
			return err
		}

		if statusInRedis != *payout.Status {
			err = s.rdbTransaction.UpdateStatus(payoutID, *payout.Status)
			if err != nil {
				return err
			}

			err = s.logsSvc.UpdateLogTransactionStatus(ctx, payout.ID, *payout.Status)
			if err != nil {
				return err
			}

			if *payout.Status == model.Succeeded || *payout.Status == model.Canceled {
				break
			}
		}

	}

	return nil
}

func signaller(ch chan<- struct{}, ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	default:
		fibArr := getFibArr()

		for _, timing := range fibArr {
			sleepTiming := time.Duration(timing) * time.Minute
			sleep(sleepTiming, ctx)
			ch <- struct{}{}
		}
		close(ch)
	}

}

func sleep(d time.Duration, ctx context.Context) {
	timer := time.NewTimer(d)

	select {
	case <-ctx.Done():
		return
	case <-timer.C:
		return
	}
}

func getFibArr() []int {
	var fibArr = []int{1, 1}

	var fibSum int
	maxFibSum := int(math.Round(redis.Expiration.Minutes()))

	index := 2

	for {
		nextFibNum := fibArr[index-1] + fibArr[index-2]
		if fibSum+nextFibNum >= maxFibSum {
			return fibArr
		}
		fibArr = append(fibArr, nextFibNum)
		index++
	}
}
