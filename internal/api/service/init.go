package service

import (
	"context"

	"github.com/ngqinzhe/ccwallet/internal/cache"
	"github.com/ngqinzhe/ccwallet/internal/db"
)

type WalletService interface {
	Deposit(ctx context.Context, req *DepositRequest) (*DepositResponse, error)
	Withdraw(ctx context.Context, req *WithdrawRequest) (*WithdrawResponse, error)
	Transfer(ctx context.Context, req *TransferRequest) (*TransferResponse, error)
	GetWalletBalance(ctx context.Context, req *GetWalletBalanceRequest) (*GetWalletBalanceResponse, error)
	GetTransactions(ctx context.Context, req *GetTransactionsRequest) (*GetTransactionsResponse, error)
}

type WalletServiceImpl struct {
	PostgreDal db.PostgreDal
	RedisCache cache.RedisCache
}

func NewWalletService(ctx context.Context, postgreDal db.PostgreDal, redisCache cache.RedisCache) WalletService {
	return &WalletServiceImpl{
		PostgreDal: postgreDal,
		RedisCache: redisCache,
	}
}
