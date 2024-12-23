package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"github.com/imperatorofdwelling/payment-svc/internal/storage/postgres"
	"go.uber.org/zap"
)

type ILogsSvc interface {
	InsertLog(ctx context.Context, log *model.Log) error
	UpdateLogStatus(ctx context.Context, payment *model.Notification) error
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

func (s *LogsSvc) UpdateLogStatus(ctx context.Context, notification *model.Notification) error {
	const op = "service.logs.UpdateLogStatus"

	notificationUUID, err := uuid.Parse(notification.Object.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = s.repo.UpdateLogStatus(ctx, notificationUUID, notification.Object.Status)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
