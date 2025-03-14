package service

import "context"

type GetAccountBalanceRequest struct {
	UserId string `json:"user_id"`
}

type GetAccountBalanceResponse struct {
	Balance map[string]float64 `json:"balance,omitempty"`
	Error   error              `json:"error,omitempty"`
}

func (a *AccountServiceImpl) GetAccountBalance(ctx context.Context, req *GetAccountBalanceRequest) (*GetAccountBalanceResponse, error) {

}
