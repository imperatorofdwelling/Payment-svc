package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
)

type ILogsRepo interface {
	InsertLog(context.Context, *model.Log) error
}

type LogsRepo struct {
	db *sql.DB
}

func NewLogsRepo(db *sql.DB) *LogsRepo {
	return &LogsRepo{db: db}
}

func (r *LogsRepo) InsertLog(ctx context.Context, log *model.Log) error {
	const op = "repo.LogsRepo.InsertLog"

	stmt, err := r.db.PrepareContext(ctx, "INSERT INTO logs(transaction_id, amount, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return fmt.Errorf("%v: %v", op, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, log.TransactionID, log.Amount, log.Status, log.CreatedAt, log.UpdatedAt)
	if err != nil {
		return fmt.Errorf("%v: %v", op, err)
	}

	return nil
}
