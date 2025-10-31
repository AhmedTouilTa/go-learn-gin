package server

import (
	"learninggo/controllers"
	"learninggo/db"
	"learninggo/middleware"

	"github.com/gin-gonic/gin"
)

func Start() {
	db.Connect()
	r := RegisterPaths()

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}

func RegisterPaths() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Authenticate())

	controllers.CheckBalance(r)
	controllers.DepositAmt(r)
	controllers.WithdrawAmt(r)
	controllers.AddUser(r)

	return r
}

func RegisterPathsNoAuth() *gin.Engine {
	r := gin.Default()

	controllers.CheckBalance(r)
	controllers.DepositAmt(r)
	controllers.WithdrawAmt(r)
	controllers.AddUser(r)

	return r
}
