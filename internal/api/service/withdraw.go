package service

import (
	"context"
	"log"
)

type WithdrawRequest struct {
	UserId            string  `json:"user_id"`
	Amount            float64 `json:"amount"`
	BankAccountNumber string  `json:"bank_account_number"`
	BankName          string  `json:"bank_name"`
}

type WithdrawResponse struct {
	Amount            float64 `json:"amount,omitempty"`
	BankAccountNumber string  `json:"bank_account_number,omitempty"`
	BankName          string  `json:"bank_name,omitempty"`
}

func (w *WalletServiceImpl) Withdraw(ctx context.Context, req *WithdrawRequest) (*WithdrawResponse, error) {
	err := w.PostgreDal.Withdraw(ctx, req.UserId, req.Amount)
	if err != nil {
		log.Printf("[AccountServiceImpl][Withdraw] failed to withdraw, dbErr: %v", err)
		return nil, err
	}
	return &WithdrawResponse{
		Amount:            req.Amount,
		BankAccountNumber: req.BankAccountNumber,
		BankName:          req.BankName,
	}, nil
}
