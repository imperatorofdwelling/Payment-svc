package repo

import (
	"database/sql"
	"go.uber.org/zap"
)

type IPaymentRepo interface {
	GetTest() string
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

func (r *PaymentRepo) GetTest() string {
	return "Hello"
}
