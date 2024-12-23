package model

type (
	Payout struct {
		Amount      `json:"amount"`
		PayoutToken string `json:"payout_token"`
		Description string `json:"description"`
		Metadata    any    `json:"metadata"`
	}
)
