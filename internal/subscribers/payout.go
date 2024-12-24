package subscribers

import (
	"fmt"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"github.com/imperatorofdwelling/payment-svc/internal/storage/redis"
)

type PayoutSubscriber struct {
	rdb redis.ITransactionRepo
}

func NewPayoutSubscriber(rdb redis.ITransactionRepo) *PayoutSubscriber {
	return &PayoutSubscriber{rdb}
}

func (p *PayoutSubscriber) Subscribe(payoutID string, status model.TransactionStatus) error {
	const op = "subscribers.payout.Subscribe"
	if status == model.Succeeded || status == model.Canceled {
		return fmt.Errorf("%s: %v", op, ErrNoNeedToCheck)
	}

	p.rdb.

	return nil
}
