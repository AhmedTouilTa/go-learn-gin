package services

import (
	"errors"
	"learninggo/db"
	"learninggo/models"
	"time"
)

func GetFirstWallet(userID string) (models.Wallet, error) {
	rows, err := db.Query("select * from wallet where user_id = $1", userID)
	if err != nil {
		return models.Wallet{}, err
	}

	for rows.Next() {
		wallet := models.Wallet{}
		err = rows.Scan(&wallet.ID, &wallet.CreationDate, &wallet.Balance, &wallet.Currency, &wallet.UserID)

		if err != nil {
			return models.Wallet{}, err
		} else {
			return wallet, nil
		}
	}

	return models.Wallet{}, errors.New("user has no wallet ")
}

func CreateWallet(userID string) error {
	_, err := db.Query("insert into wallet (creation_date,user_id,currency, balance) values($1,$2,$3,0)", time.Now(), userID, "TND")
	return err
}
