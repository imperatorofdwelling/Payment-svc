package service

import (
	"context"
	"fmt"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"github.com/imperatorofdwelling/payment-svc/internal/storage/postgres"
	"go.uber.org/zap"
	"strings"
)

type ILogsSvc interface {
	InsertLog(ctx context.Context, payment *model.Log) error
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

	event := strings.Split(notification.Event, ".")
	if len(event) != 2 {
		return fmt.Errorf("notification error in %s", op)
	}

	return nil
}
