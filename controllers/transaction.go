package controllers

import (
	"encoding/json"
	"io"
	"learninggo/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func AddTransaction(r *gin.Engine) {
	r.POST("/transaction/new", func(c *gin.Context) {
		bodyContent, _ := io.ReadAll(c.Request.Body)
		bodyData := gin.H{}
		json.Unmarshal(bodyContent, &bodyData)
		amt := bodyData["amt"].(decimal.Decimal)
		userID := strconv.FormatFloat(bodyData["user_id"].(float64), 'f', -1, 64)
		transType := bodyData["transaction_type"].(string)

		err := services.AddTransaction(userID, amt, transType)

		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "transaction added",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "adding transaction failed",
			})
		}
	})
}

func CheckBalance(r *gin.Engine) {
	r.GET("/balance/:id", func(c *gin.Context) {
		id := c.Param("id")
		_, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "id has to be a number",
			})
			return
		}

		balance, err := services.CheckBalance(id)

		if err == nil {

			c.JSON(http.StatusOK, gin.H{
				"balance": balance,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "user does not exist " + err.Error(),
			})
		}

	})

}

func DepositAmt(r *gin.Engine) {
	r.POST("/deposit", func(c *gin.Context) {
		bodContent, _ := io.ReadAll(c.Request.Body)
		bodyData := gin.H{}
		json.Unmarshal(bodContent, &bodyData)

		userID := strconv.FormatFloat(bodyData["user_id"].(float64), 'f', -1, 64)
		amt := decimal.NewFromFloat(bodyData["amt"].(float64))

		err := services.DepositAmt(userID, amt)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to deposit to balance w reason : " + err.Error(),
			})
			return
		} else {
			services.AddTransaction(userID, amt, "Deposit")
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "deposited to balance",
		})
	})
}

func WithdrawAmt(r *gin.Engine) {
	r.POST("/withdraw", func(c *gin.Context) {
		bodContent, _ := io.ReadAll(c.Request.Body)
		bodyData := gin.H{}
		json.Unmarshal(bodContent, &bodyData)

		userID := strconv.FormatFloat(bodyData["user_id"].(float64), 'f', -1, 64)
		amt := decimal.NewFromFloat(bodyData["amt"].(float64))

		err := services.WithdrawAmt(userID, amt)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to deposit to balance w reason : " + err.Error(),
			})
			return
		} else {
			services.AddTransaction(userID, amt, "Withdraw")
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "withdrew from balance",
		})
	})
}

/* func PerformTransaction(r *gin.Engine) {
	r.POST("/perform_transaction", func(c *gin.Context) {
		primaryUserId := c.PostForm("primary_user_id")
		secondaryUserId := c.PostForm("_user_id")

		if transactionType == ""
	})
} */
