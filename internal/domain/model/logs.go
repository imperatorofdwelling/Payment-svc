package model

import (
	"github.com/google/uuid"
	"time"
)

type LogPaymentRequest struct {
	TransactionID uuid.UUID `json:"transaction_id" validate:"required,uuid"`
	Status        string    `json:"status"`
	Value         string    `json:"value" validate:"required,money"`
	Currency      string    `json:"currency" validate:"required,iso4217"`
	Type          string    `json:"type" validate:"required"`
}

type LogPayment struct {
	ID            uuid.UUID `json:"id"`
	TransactionID uuid.UUID `json:"transaction_id" validate:"required,uuid"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type LogPaymentAmount struct {
	ID        uuid.UUID `json:"id"`
	Value     string    `json:"value" validate:"required,money"`
	Currency  string    `json:"currency" validate:"required,iso4217"`
	PaymentID uuid.UUID `json:"payment_id" validate:"required,uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LogPaymentMethod struct {
	ID        uuid.UUID `json:"id"`
	PaymentID uuid.UUID `json:"payment_id" validate:"required,uuid"`
	Type      string    `json:"type" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
