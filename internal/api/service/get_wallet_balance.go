package service

import (
	"context"
	"log"
)

type GetWalletBalanceRequest struct {
	UserId string `json:"user_id"`
}

type GetWalletBalanceResponse struct {
	Balance float64 `json:"balance,omitempty"`
}

func (w *WalletServiceImpl) GetWalletBalance(ctx context.Context, req *GetWalletBalanceRequest) (*GetWalletBalanceResponse, error) {
	// get from cache first
	value, err := w.RedisCache.Get(ctx, req.UserId)
	if err == nil {
		log.Printf("[WalletServiceImpl][GetWalletBalance] retrieved balance: %f from cache for user: %s", value.(float64), req.UserId)
		return &GetWalletBalanceResponse{
			Balance: value.(float64),
		}, nil
	}

	// get from db
	balance, err := w.PostgreDal.GetWalletBalance(ctx, req.UserId)
	if err != nil {
		log.Printf("[WalletServiceImpl][GetWalletBalance] failed to get wallet balance, dbErr: %v", err)
		return nil, err
	}
	if err := w.RedisCache.Set(ctx, req.UserId, balance); err != nil {
		log.Printf("[WalletServiceImpl][GetWalletBalance] failed to set to redisCache, err: %v", err)
	}
	log.Printf("[WalletServiceImpl][GetWalletBalance] successfully getWalletBalance for user: %s, balance: %f", req.UserId, balance)
	return &GetWalletBalanceResponse{
		Balance: balance,
	}, nil
}
