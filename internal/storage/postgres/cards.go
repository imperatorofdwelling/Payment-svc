package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"go.uber.org/zap"
	"time"
)

type ICardsRepo interface {
	CreateCard(context.Context, model.Card) error
	UpdateCard(context.Context, model.Card) error
	CardSynonymIsExists(ctx context.Context, synonym string) (bool, error)
	CheckCardExistsByID(ctx context.Context, cardID uuid.UUID) (bool, error)
	DeleteCardByID(context.Context, uuid.UUID) error
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

	stmt, err := r.db.PrepareContext(ctx, "INSERT INTO bank_cards(user_id, bank_name, country_code, synonim, card_mask, type, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		return fmt.Errorf("%v: %v", op, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, card.UserId, card.BankName, card.CountryCode, card.Synonim, card.CardMask, card.Type, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("%v: %v", op, err)
	}

	return nil
}

func (r *CardsRepo) UpdateCard(ctx context.Context, card model.Card) error {
	const op = "repo.postgres.card.UpdateCard"

	stmt, err := r.db.PrepareContext(ctx, "UPDATE bank_cards SET bank_name=$1, country_code=$2, synonim=$3, card_mask=$4, type=$5, updated_at=$6 WHERE user_id=$7")
	if err != nil {
		return fmt.Errorf("%v: %v", op, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, card.BankName, card.CountryCode, card.Synonim, card.CardMask, card.Type, time.Now(), card.UserId)
	if err != nil {
		return fmt.Errorf("%v: %v", op, err)
	}

	return nil
}

func (r *CardsRepo) CardSynonymIsExists(ctx context.Context, synonym string) (bool, error) {
	const op = "repo.postgres.card.CardSynonymIsExists"

	stmt, err := r.db.PrepareContext(ctx, "SELECT * FROM bank_cards WHERE synonim = $1 LIMIT 1")
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

func (r *CardsRepo) CheckCardExistsByID(ctx context.Context, cardID uuid.UUID) (bool, error) {
	const op = "repo.postgres.card.CheckCardExistsByID"

	stmt, err := r.db.PrepareContext(ctx, "SELECT EXISTS(SELECT * FROM bank_cards WHERE id = $1)")
	if err != nil {
		return false, fmt.Errorf("%v: %v", op, err)
	}

	defer stmt.Close()

	var exists bool

	err = stmt.QueryRowContext(ctx, cardID).Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("%v: %v", op, err)
	}

	return exists, nil
}

func (r *CardsRepo) DeleteCardByID(ctx context.Context, cardID uuid.UUID) error {
	const op = "repo.postgres.card.DeleteCardByID"

	stmt, err := r.db.PrepareContext(ctx, "DELETE FROM bank_cards WHERE id = $1")
	if err != nil {
		return fmt.Errorf("%v: %v", op, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, cardID)
	if err != nil {
		return fmt.Errorf("%v: %v", op, err)
	}

	return nil
}
