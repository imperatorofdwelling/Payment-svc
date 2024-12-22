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
	InsertLog(context.Context, *model.Log) error
	CheckTransactionIDExists(ctx context.Context, transactionID uuid.UUID) (bool, error)
	UpdateLogStatus(ctx context.Context, transactionID uuid.UUID, status model.TransactionStatus) error
}

type LogsRepo struct {
	db  *sql.DB
	log *zap.SugaredLogger
}

func NewLogsRepo(db *sql.DB, log *zap.SugaredLogger) *LogsRepo {
	return &LogsRepo{db, log}
}

func (r *LogsRepo) InsertLog(ctx context.Context, p *model.Log) error {
	const op = "repo.postgres.logs.InsertLog"

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer tx.Rollback()

	isExist, err := r.CheckTransactionIDExists(ctx, p.TransactionID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if isExist {
		return nil
	}

	stmtPayment, err := tx.PrepareContext(ctx, "INSERT INTO logs(transaction_id, method_type, transaction_type, status, value, currency, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer stmtPayment.Close()

	_, err = stmtPayment.ExecContext(ctx, p.TransactionID, p.MethodType, p.TransactionType, p.Status, p.Value, p.Currency, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%v: %v", op, err)
	}

	return nil
}

func (r *LogsRepo) CheckTransactionIDExists(ctx context.Context, transactionID uuid.UUID) (bool, error) {
	const op = "repo.postgres.logs.CheckTransactionIDExists"

	stmt, err := r.db.PrepareContext(ctx, "SELECT EXISTS(SELECT * FROM logs WHERE transaction_id = $1)")
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	var exists bool

	err = stmt.QueryRowContext(ctx, transactionID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return exists, nil
}

func (r *LogsRepo) UpdateLogStatus(ctx context.Context, transactionID uuid.UUID, status model.TransactionStatus) error {
	const op = "repo.postgres.logs.UpdateLogStatus"

	stmt, err := r.db.PrepareContext(ctx, `UPDATE logs SET status = $1 WHERE transaction_id = $2`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, status, transactionID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
