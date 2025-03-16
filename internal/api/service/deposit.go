package service

import (
	"context"
	"log"
)

type DepositRequest struct {
	UserId string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

type DepositResponse struct {
	Amount float64 `json:"amount,omitempty"`
}

func (w *WalletServiceImpl) Deposit(ctx context.Context, req *DepositRequest) (*DepositResponse, error) {
	err := w.PostgreDal.Deposit(ctx, req.UserId, req.Amount)
	if err != nil {
		log.Printf("[WalletServiceImpl][Deposit] deposit failed, dbErr: %v", err)
		return nil, err
	}
	return &DepositResponse{
		Amount: req.Amount,
	}, nil
}
