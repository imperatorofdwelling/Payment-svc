package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"go.uber.org/zap"
	"time"
)

type ICardsRepo interface {
	CreateCard(context.Context, model.Card) error
	CardSynonymIsExists(ctx context.Context, synonym string) (bool, error)
}

type CardsRepo struct {
	db  *sql.DB
	log *zap.SugaredLogger
}

func NewCardsRepo(db *sql.DB, log *zap.SugaredLogger) *CardsRepo {
	return &CardsRepo{db: db, log: log}
}

func (r *CardsRepo) CreateCard(ctx context.Context, card model.Card) error {
	const op = "repo.postgres.card.CreateCard"

	stmt, err := r.db.PrepareContext(ctx, "INSERT INTO bank_cards(user_id, bank_name, country_code, synonym, card_mask, type, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		return fmt.Errorf("%v: %v", op, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, card.UserId, card.BankName, card.CountryCode, card.Synonym, card.CardMask, card.Type, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("%v: %v", op, err)
	}

	return nil
}

func (r *CardsRepo) CardSynonymIsExists(ctx context.Context, synonym string) (bool, error) {
	const op = "repo.postgres.card.CardSynonymIsExists"

	stmt, err := r.db.PrepareContext(ctx, "SELECT * FROM bank_cards WHERE synonym = $1 LIMIT 1")
	if err != nil {
		return false, fmt.Errorf("%v: %v", op, err)
	}

	defer stmt.Close()

	var synonymID string

	err = stmt.QueryRowContext(ctx, synonym).Scan(&synonymID)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("%v: %v", op, err)
	}

	return true, nil
}
