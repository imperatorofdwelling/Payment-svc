package model

import (
	"github.com/google/uuid"
	"time"
)

type Log struct {
	ID              uuid.UUID         `json:"id"`
	TransactionID   uuid.UUID         `json:"transaction_id" validate:"required,uuid"`
	TransactionType TransactionType   `json:"transaction_type" validate:"required"`
	Status          TransactionStatus `json:"status" validate:"required"`
	Value           string            `json:"value" validate:"required,money"`
	Currency        Currency          `json:"currency" validate:"required,iso4217"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
}
