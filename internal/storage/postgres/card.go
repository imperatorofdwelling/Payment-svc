package postgres

import "database/sql"

type ICardRepo interface {
}

type CardRepo struct {
	db *sql.DB
}

func NewCardRepo(db *sql.DB) *CardRepo {
	return &CardRepo{db: db}
}
