package model

import (
	"time"

	usr "github.com/muosilva/bankgo/internal"
)

type Account struct {
	ID        uint
	AccNumber string
	User      usr.User
	Balance   float64
	Currency  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
