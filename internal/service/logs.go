package service

import (
	"context"
	"fmt"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"github.com/imperatorofdwelling/payment-svc/internal/storage/postgres"
	"go.uber.org/zap"
)

type ILogsSvc interface {
	InsertLog(ctx context.Context, log *model.Log) error
	UpdateLogTransactionStatus(ctx context.Context, transactionId string, status model.TransactionStatus) error
}

type LogsSvc struct {
	repo postgres.ILogsRepo
	log  *zap.SugaredLogger
}

func NewLogsService(repo postgres.ILogsRepo, log *zap.SugaredLogger) *LogsSvc {
	return &LogsSvc{repo, log}
}

func (s *LogsSvc) InsertLog(ctx context.Context, log *model.Log) error {
	const op = "service.logs.InsertLog"

	err := s.repo.InsertLog(ctx, log)
	if err != nil {
		return err
	}

	return nil
}

func (s *LogsSvc) UpdateLogTransactionStatus(ctx context.Context, transactionId string, status model.TransactionStatus) error {
	const op = "service.logs.UpdateLogStatus"

	err := s.repo.UpdateLogStatus(ctx, transactionId, status)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
