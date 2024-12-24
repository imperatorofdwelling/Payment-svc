package redis

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"github.com/redis/go-redis/v9"
	"time"
)

type ITransactionRepo interface {
	Commit() error
	UpdateStatus() error
	GetStatus() error
	IsExists(id uuid.UUID) (bool, error)
}

var (
	ErrTransactionAlreadyExists = errors.New("transaction already exists")
	ErrTransactionNotFound      = errors.New("transaction not found")
	ErrChangedKeyErr            = errors.New("the key changed at the time of the request")
)

const (
	TransactionTable = "transactionTable"
	expiration       = time.Minute * 24 * 60
)

type TransactionRepo struct {
	rdb *redis.Client
}

func NewTransactionRepo(rdb *redis.Client) *TransactionRepo {
	return &TransactionRepo{rdb: rdb}
}

func (r *TransactionRepo) Commit(id uuid.UUID, status model.TransactionStatus) error {
	const op = "redis.transaction.Commit"

	exists, err := r.IsExists(id)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("%s: %w", op, ErrTransactionAlreadyExists)
	}

	pipe := r.rdb.TxPipeline()
	pipe.Set(ctx, r.getKey(id), status, expiration)
	_, err = pipe.Exec(ctx)

	return err
}

func (r *TransactionRepo) UpdateStatus(id uuid.UUID, status model.TransactionStatus) error {
	const op = "redis.transaction.UpdateStatus"

	exists, err := r.IsExists(id)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("%s: %w", op, ErrTransactionNotFound)
	}

	ttl, err := r.rdb.TTL(ctx, r.getKey(id)).Result()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	duration := time.Since(time.Unix(int64(ttl), 0))

	pipe := r.rdb.TxPipeline()
	pipe.Set(ctx, r.getKey(id), status, duration)
	_, err = pipe.Exec(ctx)
	return err
}

func (r *TransactionRepo) GetStatus(id uuid.UUID) (model.TransactionStatus, error) {
	const op = "redis.transaction.GetStatus"

	exists, err := r.IsExists(id)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if !exists {
		return "", fmt.Errorf("%s: %w", op, ErrTransactionNotFound)
	}

	transactionKey := r.getKey(id)
	var transactionStatus model.TransactionStatus

	err = r.rdb.Watch(ctx, func(tx *redis.Tx) error {
		status, err := tx.Get(ctx, transactionKey).Result()
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		transactionStatus = model.TransactionStatus(status)
		return nil
	}, transactionKey)
	if errors.Is(err, redis.TxFailedErr) {
		return "", fmt.Errorf("%s: %w", op, ErrChangedKeyErr)
	}
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if transactionStatus == model.Succeeded || transactionStatus == model.Canceled {
		defer r.delKey(id)
	}

	return transactionStatus, nil
}

func (r *TransactionRepo) IsExists(id uuid.UUID) (bool, error) {
	const op = "redis.transaction.IsExists"

	val, err := r.rdb.Exists(ctx, r.getKey(id)).Result()
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return val == 1, nil
}

func (r *TransactionRepo) getKey(id uuid.UUID) string {
	return TransactionTable + ":" + id.String()
}

func (r *TransactionRepo) delKey(id uuid.UUID) error {
	return r.rdb.Del(ctx, r.getKey(id)).Err()
}
