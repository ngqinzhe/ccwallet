package handler

import (
	"context"

	"github.com/ngqinzhe/ccwallet/internal/api/service"
	"github.com/ngqinzhe/ccwallet/internal/db"
)

type WalletController struct {
	WalletService service.WalletService
}

func NewWalletController(ctx context.Context, db db.PostgreDal) *WalletController {
	return &WalletController{
		WalletService: service.NewWalletService(ctx, db),
	}
}
