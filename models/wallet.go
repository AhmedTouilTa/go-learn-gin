package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Wallet struct {
	ID           int             `json:"id"`
	CreationDate time.Time       `json:"creation_time"`
	Balance      decimal.Decimal `json:"balance"`
	Currency     string          `json:"currency"`
	UserID       int             `json:"user_id"`
}
