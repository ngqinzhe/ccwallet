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
	Transactions []*model.Transaction `json:"transactions,omitempty"`
}

func (a *AccountServiceImpl) GetTransactions(ctx context.Context, req *GetTransactionHistoryRequest) (*GetTransactionHistoryResponse, error) {

}
