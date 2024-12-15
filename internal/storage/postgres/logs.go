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

type ILogsRepo interface {
	InsertLog(context.Context, *model.LogPaymentRequest) error
}

type LogsRepo struct {
	db  *sql.DB
	log *zap.SugaredLogger
}

func NewLogsRepo(db *sql.DB, log *zap.SugaredLogger) *LogsRepo {
	return &LogsRepo{db, log}
}

func (r *LogsRepo) InsertLog(ctx context.Context, p *model.LogPaymentRequest) error {
	const op = "repo.postgres.logs.InsertLog"

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer tx.Rollback()

	stmtPayment, err := tx.PrepareContext(ctx, "INSERT INTO payment_logs(id, transaction_id, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer stmtPayment.Close()

	paymentID := uuid.New()

	_, err = stmtPayment.ExecContext(ctx, paymentID, p.TransactionID, p.Status, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	stmtAmount, err := tx.PrepareContext(ctx, "INSERT INTO payment_logs_amount(value, currency, payment_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer stmtAmount.Close()

	_, err = stmtAmount.ExecContext(ctx, p.Value, p.Currency, paymentID, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	stmtMethod, err := tx.PrepareContext(ctx, "INSERT INTO payment_logs_method(payment_id, type, created_at, updated_at) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer stmtMethod.Close()

	_, err = stmtMethod.ExecContext(ctx, paymentID, p.Type, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%v: %v", op, err)
	}

	return nil
}
