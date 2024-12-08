package model

const (
	CardLen int = 16
)

type Card struct {
	ID          int    `json:"id"`
	UserId      string `json:"user_id"`
	CardSynonym string `json:"card_synonym"`
	CardMask    string `json:"card_mask"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}
