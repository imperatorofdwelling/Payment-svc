package model

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Balance   float32   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BalanceChange struct {
}

func (u *User) ToString() string {
	return fmt.Sprintf("ID: %s, Balance: %f", u.ID, u.Balance)
}
