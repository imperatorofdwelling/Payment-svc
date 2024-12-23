package postgres

import (
	"database/sql"
	"go.uber.org/zap"
)

type IPayoutsRepo interface {
}

type PayoutsRepo struct {
	db  *sql.DB
	log *zap.SugaredLogger
}

func NewPayoutsRepo(db *sql.DB, log *zap.SugaredLogger) *PayoutsRepo {
	return &PayoutsRepo{db, log}
}
