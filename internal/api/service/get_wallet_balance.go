package service

import "context"

type GetWalletBalanceRequest struct {
	UserId string `json:"user_id"`
}

type GetWalletBalanceResponse struct {
	Balance map[string]float64 `json:"balance,omitempty"`
	Error   error              `json:"error,omitempty"`
}

func (a *AccountServiceImpl) GetWalletBalance(ctx context.Context, req *GetWalletBalanceRequest) (*GetWalletBalanceResponse, error) {

}
