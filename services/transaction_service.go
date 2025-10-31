package services

import (
	"fmt"
	"learninggo/db"
	"time"

	"github.com/shopspring/decimal"
)

func AddTransaction(userID string, amt decimal.Decimal, transactionType string) error {
	wallet, werr := GetFirstWallet(userID)

	if werr != nil {
		fmt.Println(werr.Error())
	}

	_, err := db.Query("insert into transaction (primary_wallet_id,amount,transaction_type,transaction_date) values ($1,$2,$3,$4)", wallet.ID, amt, transactionType, time.Now())

	return err
}
