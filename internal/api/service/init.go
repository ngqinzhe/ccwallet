package service

import (
	"context"

	"github.com/ngqinzhe/ccwallet/internal/db"
)

type AccountService interface {
	Deposit(ctx context.Context, req *DepositRequest) (*DepositResponse, error)
	Withdraw(ctx context.Context, req *WithdrawRequest) (*WithdrawResponse, error)
	Transfer(ctx context.Context, req *TransferRequest) (*TransferResponse, error)
	GetWalletBalance(ctx context.Context, req *GetWalletBalanceRequest) (*GetWalletBalanceResponse, error)
	GetTransactions(ctx context.Context, req *GetTransactionHistoryRequest) (*GetTransactionHistoryResponse, error)
}

type AccountServiceImpl struct {
	PostgreDal db.PostgreDal
}

func NewAccountService(ctx context.Context, postgreDal db.PostgreDal) AccountService {
	return &AccountServiceImpl{
		PostgreDal: postgreDal,
	}
}
