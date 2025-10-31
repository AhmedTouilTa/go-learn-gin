package test

import (
	"encoding/json"
	"fmt"
	"learninggo/db"
	"learninggo/models"
	"learninggo/server"
	"learninggo/services"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/go-jose/go-jose/v4/testutils/assert"
	"github.com/shopspring/decimal"
)

func TestAddUserRoute(t *testing.T) {
	db.Connect()
	router := server.RegisterPathsNoAuth()

	w := httptest.NewRecorder()
	user := models.User{
		Username: "kalb2",
	}

	userJson, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/user/new", strings.NewReader(string(userJson)))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
	//assert.Equal(t, "balance", w.Body.String())
}

func TestBalanceRoute(t *testing.T) {
	db.Connect()
	router := server.RegisterPathsNoAuth()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/balance/0", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	//assert.Equal(t, "balance", w.Body.String())
}

func TestAddUserService(t *testing.T) {
	db.Connect()
	err := services.AddUser("newUser1")
	assert.NoError(t, err)
}

func TestDespositAmtService(t *testing.T) {
	db.Connect()
	useri, _ := services.GetFirstUser()
	err := services.DepositAmt(strconv.FormatInt(int64(useri.ID), 10), decimal.NewFromInt(1050))
	assert.NoError(t, err)
}

func TestWithdrawAmtService(t *testing.T) {
	db.Connect()
	useri, _ := services.GetFirstUser()
	err := services.WithdrawAmt(strconv.FormatInt(int64(useri.ID), 10), decimal.NewFromInt(1050))
	assert.NoError(t, err)
}

func TestCheckBalanceService(t *testing.T) {
	db.Connect()
	useri, _ := services.GetFirstUser()
	_, err := services.CheckBalance(strconv.FormatInt(int64(useri.ID), 10))
	assert.NoError(t, err)
}

func TestGetFirstWallet(t *testing.T) {
	db.Connect()
	useri, _ := services.GetFirstUser()
	_, err := services.GetFirstWallet(strconv.FormatInt(int64(useri.ID), 10))
	assert.NoError(t, err)
}
