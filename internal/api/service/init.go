package service

import (
	"context"
)

type AccountServiceInf interface {
	Deposit(ctx context.Context, req *DepositRequest) *DepositResponse
	Withdraw(ctx context.Context, req *WithdrawRequest) *WithdrawResponse
	Transfer(ctx context.Context, req *TransferRequest) *TransferResponse
	GetAccountBalance(ctx context.Context, req *GetAccountBalanceRequest) *GetAccountBalanceResponse
	GetTransactionHistory(ctx context.Context, req *GetTransactionHistoryRequest) *GetTransactionHistoryResponse
}

type AccountService struct {
}

func NewAccountService(ctx context.Context) AccountServiceInf {
	return &AccountService{}
}
