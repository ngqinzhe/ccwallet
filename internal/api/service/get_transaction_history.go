package service

import (
	"context"
	"time"

	"github.com/ngqinzhe/ccwallet/internal/model"
)

type GetTransactionHistoryRequest struct {
	UserId   string    `json:"user_id"`
	FromDate time.Time `json:"from_date,omitempty"`
	ToDate   time.Time `json:"to_date,omitempty"`
}

type GetTransactionHistoryResponse struct {
	Events []*model.Event `json:"events,omitempty"`
	Error  error          `json:"error"`
}

func (a *AccountServiceImpl) GetTransactionHistory(ctx context.Context, req *GetTransactionHistoryRequest) (*GetTransactionHistoryResponse, error) {

}
