package service

import "context"

type TransferRequest struct {
	FromUserId int64   `json:"from_user_id"`
	ToUserId   int64   `json:"to_user_id"`
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
}

type TransferResponse struct {
	Amount   float64 `json:"amount,omitempty"`
	Currency string  `json:"currency,omitempty"`
	Error    error   `json:"error,omitempty"`
}

func (a *AccountService) Transfer(ctx context.Context, req *TransferRequest) *TransferResponse {

}
