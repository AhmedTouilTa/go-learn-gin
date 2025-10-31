package services

import (
	"errors"
	"fmt"
	"learninggo/db"
	"learninggo/models"
	"log"
	"strconv"

	"github.com/shopspring/decimal"
)

func CheckBalance(id string) (decimal.Decimal, error) {

	rows, err := db.Query("select balance from wallet where user_id = $1", id)

	if err != nil {
		log.Fatal(err)
	}

	balance := decimal.NewFromInt(0)
	if rows.Next() {
		err := rows.Scan(&balance)
		if err != nil {
			log.Fatal(err)
		}
		return balance, err
	} else {
		return decimal.Decimal{}, err
	}
}

func AddUser(name string) error {
	rows, err := db.Query("insert into public.user (username) values ($1) returning id", name)
	if err != nil {
		fmt.Println(err.Error())
	}
	var newUserID uint
	for rows.Next() {
		rows.Scan(&newUserID)
	}
	fmt.Println(newUserID)
	err = CreateWallet(strconv.FormatUint(uint64(newUserID), 10))

	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func DepositAmt(userID string, amt decimal.Decimal) error {
	_, err := db.Query("update wallet set balance = balance + $1 where user_id=$2", amt, userID)
	return err
}

func WithdrawAmt(userID string, amt decimal.Decimal) error {
	_, err := db.Query("update wallet set balance = balance - $1 where user_id=$2", amt, userID)
	return err
}

func GetFirstUser() (models.User, error) {
	rows, err := db.Query("select * from public.user")
	if err != nil {
		return models.User{}, err
	}

	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.ID, &user.Username)

		if err != nil {
			return models.User{}, err
		} else {
			return user, nil
		}
	}

	return models.User{}, errors.New(" ")
}
