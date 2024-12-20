package model

import (
	"github.com/google/uuid"
	"time"
)

type Card struct {
	ID          uuid.UUID `json:"id"`
	UserId      string    `json:"user_id"`
	BankName    string    `json:"bank_name,omitempty"`
	CountryCode string    `json:"country_code"`
	Synonym     string    `json:"synonym"`
	CardMask    string    `json:"card_mask"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
