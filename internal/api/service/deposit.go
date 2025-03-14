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

func (a *AccountServiceImpl) Deposit(ctx context.Context, req *DepositRequest) (*DepositResponse, error) {
	err := a.PostgreDal.Deposit(ctx, req.UserId, req.Amount)
	if err != nil {
		log.Printf("[AccountServiceImpl][Deposit] deposit failed, dbErr: %v", err)
		return nil, err
	}
	return &DepositResponse{
		Amount: req.Amount,
	}, nil
}
