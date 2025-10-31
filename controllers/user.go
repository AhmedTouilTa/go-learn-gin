package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"learninggo/models"
	"learninggo/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUser(r *gin.Engine) {
	r.POST("/user/new", func(c *gin.Context) {
		bodContent, _ := io.ReadAll(c.Request.Body)
		user := models.User{}
		json.Unmarshal(bodContent, &user)

		err := services.AddUser(user.Username)
		if err != nil {
			fmt.Println("failed to add user because " + err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "added user of name " + user.Username,
		})
	})
}
