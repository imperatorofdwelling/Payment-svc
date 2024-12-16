package service

import (
	"context"
	"fmt"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"github.com/imperatorofdwelling/payment-svc/internal/storage/postgres"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ICardsSvc interface {
	CreateBankCard(ctx context.Context, card model.Card) error
}

type CardsSvc struct {
	repo postgres.ICardsRepo
	log  *zap.SugaredLogger
}

func NewCardsService(repo postgres.ICardsRepo, log *zap.SugaredLogger) *CardsSvc {
	return &CardsSvc{repo, log}
}

func (s *CardsSvc) CreateBankCard(ctx context.Context, card model.Card) error {
	const op = "service.cards.CreateCard"

	isExists, err := s.repo.CardSynonymIsExists(ctx, card.Synonym)
	if err != nil {
		return errors.Wrap(err, op)
	}

	if isExists {
		return fmt.Errorf("%s: %v", op, ErrCardAlreadyExists)
	}

	err = s.repo.CreateCard(ctx, card)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}
