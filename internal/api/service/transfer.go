package service

import (
	"context"
	"log"
)

type TransferRequest struct {
	FromUserId string  `json:"from_user_id"`
	ToUserId   string  `json:"to_user_id"`
	Amount     float64 `json:"amount"`
}

type TransferResponse struct {
	Amount float64 `json:"amount,omitempty"`
}

func (w *WalletServiceImpl) Transfer(ctx context.Context, req *TransferRequest) (*TransferResponse, error) {
	err := w.PostgreDal.Transfer(ctx, req.FromUserId, req.ToUserId, req.Amount)
	if err != nil {
		log.Printf("[WalletServiceImpl][Transfer] transfer failed, dbErr: %v", err)
		return nil, err
	}
	return &TransferResponse{
		Amount: req.Amount,
	}, nil
}
