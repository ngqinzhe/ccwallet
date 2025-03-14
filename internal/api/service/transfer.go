package service

import "context"

type TransferRequest struct {
	FromUserId string  `json:"from_user_id"`
	ToUserId   string  `json:"to_user_id"`
	Amount     float64 `json:"amount"`
}

type TransferResponse struct {
	Amount float64 `json:"amount,omitempty"`
}

func (a *AccountServiceImpl) Transfer(ctx context.Context, req *TransferRequest) (*TransferResponse, error) {

}
