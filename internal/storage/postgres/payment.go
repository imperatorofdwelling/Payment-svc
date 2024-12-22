package postgres

import (
	"database/sql"
	"go.uber.org/zap"
)

type IPaymentRepo interface {
}

type PaymentRepo struct {
	db  *sql.DB
	log *zap.SugaredLogger
}

func NewPaymentRepo(db *sql.DB, log *zap.SugaredLogger) *PaymentRepo {
	return &PaymentRepo{
		db:  db,
		log: log,
	}
}
