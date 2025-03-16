package main

import (
	"context"

	"github.com/ngqinzhe/ccwallet/internal/api/handler"
	"github.com/ngqinzhe/ccwallet/internal/api/router"
	"github.com/ngqinzhe/ccwallet/internal/db"
	"github.com/ngqinzhe/ccwallet/internal/util"
)

func main() {
	ctx := context.Background()
	config := util.InitConfig()
	dal := db.NewPostgreDal(ctx, &config.PostgreSqlCredentials)

	walletController := handler.NewWalletController(ctx, dal)

	router.Run(ctx, walletController)
}
