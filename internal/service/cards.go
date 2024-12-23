package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"github.com/imperatorofdwelling/payment-svc/internal/storage/postgres"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ICardsSvc interface {
	CreateBankCard(ctx context.Context, card model.Card) error
	DeleteCardByID(ctx context.Context, cardID uuid.UUID) error
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

	isExists, err := s.repo.CardSynonymIsExists(ctx, card.Synonim)
	if err != nil {
		return errors.Wrap(err, op)
	}

	if isExists {
		return fmt.Errorf("%s: %v", op, ErrCardAlreadyExists)
	}

	userIDExists, err := s.repo.CheckCardExistsByID(ctx, card.ID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	if userIDExists {
		err = s.repo.UpdateCard(ctx, card)
		if err != nil {
			return errors.Wrap(err, op)
		}
	} else {
		err = s.repo.CreateCard(ctx, card)
		if err != nil {
			return errors.Wrap(err, op)
		}
	}

	return nil
}

func (s *CardsSvc) DeleteCardByID(ctx context.Context, cardID uuid.UUID) error {
	const op = "service.cards.DeleteCardByID"

	userIDExists, err := s.repo.CheckCardExistsByID(ctx, cardID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	if !userIDExists {
		return fmt.Errorf("card  with user id %v not found", cardID)
	}

	err = s.repo.DeleteCardByID(ctx, cardID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}
