package service

import "context"

type DepositRequest struct {
	UserId   int64   `json:"user_id"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type DepositResponse struct {
	Amount   float64 `json:"amount,omitempty"`
	Currency string  `json:"currency,omitempty"`
	Error    error   `json:"error,omitempty"`
}

func (a *AccountService) Deposit(ctx context.Context, req *DepositRequest) *DepositResponse {

}
