package router

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ngqinzhe/ccwallet/internal/api/handler"
)

func Run(ctx context.Context) {
	router := gin.Default()
	router.Use(cors.Default())

	// Writes
	router.POST("/api/v1/deposit", handler.Deposit(ctx))
	router.POST("/api/v1/withdraw", handler.Withdraw(ctx))
	router.POST("/api/v1/transfer", handler.Withdraw(ctx))

	// Reads
	router.GET("/api/v1/account_balance", handler.GetAccountBalance(ctx))
	router.GET("/api/v1/transaction_history", handler.GetTransactionHistory(ctx))

	router.Run("localhost:8080")

	// router.Run(env.EnvConfig.ServerEndpoint)
}
