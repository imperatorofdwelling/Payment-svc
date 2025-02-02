package model

import (
	yoomodel "github.com/eclipsemode/go-yookassa-sdk/yookassa/model"
	"github.com/google/uuid"
	"time"
)

type Log struct {
	ID              uuid.UUID                  `json:"id"`
	TransactionID   string                     `json:"transaction_id" validate:"required"`
	TransactionType yoomodel.TransactionType   `json:"transaction_type" validate:"required"`
	Status          yoomodel.TransactionStatus `json:"status" validate:"required"`
	Value           string                     `json:"value" validate:"required,money"`
	Currency        yoomodel.Currency          `json:"currency" validate:"required,iso4217"`
	CreatedAt       time.Time                  `json:"created_at"`
	UpdatedAt       time.Time                  `json:"updated_at"`
}
