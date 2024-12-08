package redis

import (
	"errors"
	"github.com/redis/go-redis/v9"
)

type ITransactionRepo interface {
	Commit() error
	UpdateStatus() error
	GetStatus() error
	IsExists() (bool, error)
}

var (
	TransactionAlreadyExistsErr = errors.New("transactionAlreadyExistsError")
	TransactionNotFoundErr      = errors.New("transactionNotFoundError")
	ChangedKeyErr               = errors.New("the key changed at the time of the request")
)

type TransactionRepo struct {
	rdb *redis.Client
}

func NewTransactionRepo(rdb *redis.Client) *TransactionRepo {
	return &TransactionRepo{rdb: rdb}
}

func (r *TransactionRepo) Commit() error {
	panic("implement me")
}

func (r *TransactionRepo) UpdateStatus() error {
	panic("implement me")
}

func (r *TransactionRepo) GetStatus() error {
	panic("implement me")
}

func (r *TransactionRepo) IsExists() (bool, error) {
	panic("implement me")
}
