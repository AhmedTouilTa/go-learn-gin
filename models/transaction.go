package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID                int             `json:"id"`
	PrimaryWalletID   int             `json:"primary_wallet_id"`
	RecipientWalletID int             `json:"recepient_wallet_id"`
	TransactionType   string          `json:"transaction_type"`
	Amount            decimal.Decimal `json:"amount"`
	TransactionDate   time.Time       `json:"transaction_date"`
}
