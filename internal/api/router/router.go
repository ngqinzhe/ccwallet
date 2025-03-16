package router

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ngqinzhe/ccwallet/internal/api/handler"
)

func Run(ctx context.Context, walletController *handler.WalletController) {
	router := gin.Default()
	router.Use(cors.Default())

	// Writes
	router.POST("/api/v1/deposit", walletController.Deposit(ctx))
	router.POST("/api/v1/withdraw", walletController.Withdraw(ctx))
	router.POST("/api/v1/transfer", walletController.Withdraw(ctx))

	// Reads
	router.GET("/api/v1/wallet_balance", walletController.GetWalletBalance(ctx))
	router.GET("/api/v1/transactions", walletController.GetTransactions(ctx))

	router.Run("localhost:8080")

	// router.Run(env.EnvConfig.ServerEndpoint)
}
