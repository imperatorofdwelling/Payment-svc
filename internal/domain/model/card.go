package model

import (
	"github.com/google/uuid"
	"time"
)

type Card struct {
	ID            uuid.UUID `json:"id"`
	UserId        uuid.UUID `json:"user_id"`
	IssuerCountry string    `json:"issuer_country,omitempty" validate:"omitempty,iso3166_1_alpha2"`
	PayoutToken   string    `json:"payout_token"`
	First6        string    `json:"first6" validate:"required,numeric"`
	Last4         string    `json:"last4" validate:"required,numeric"`
	CardType      string    `json:"card_type"`
	IssuerName    string    `json:"issuer_name,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
