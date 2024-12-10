package model

import "time"

const (
	CardLen int = 16
)

type Card struct {
	ID          int       `json:"id"`
	UserId      string    `json:"user_id"`
	CardSynonym string    `json:"card_synonym"`
	CardMask    string    `json:"card_mask"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
