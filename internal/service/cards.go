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
	DeleteCardByUserID(ctx context.Context, userID uuid.UUID) error
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

	userIDUUID, err := uuid.Parse(card.UserId)
	if err != nil {
		return errors.Wrap(err, op)
	}

	userIDExists, err := s.repo.CheckCardExistsByUserID(ctx, userIDUUID)
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

func (s *CardsSvc) DeleteCardByUserID(ctx context.Context, userID uuid.UUID) error {
	const op = "service.cards.DeleteCardByUserID"

	userIDExists, err := s.repo.CheckCardExistsByUserID(ctx, userID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	if !userIDExists {
		return fmt.Errorf("card  with user id %v not found", userID)
	}

	err = s.repo.DeleteCardByUserID(ctx, userID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}
