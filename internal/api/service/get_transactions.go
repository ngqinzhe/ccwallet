package service

import (
	"context"
	"log"
	"time"

	"github.com/ngqinzhe/ccwallet/internal/model"
)

type GetTransactionsRequest struct {
	UserId   string    `json:"user_id"`
	FromDate time.Time `json:"from_date,omitempty"`
	ToDate   time.Time `json:"to_date,omitempty"`
	Limit    *int      `json:"limit,omitempty"`
	Offset   *int      `json:"limit,omitempty"`
}

type GetTransactionsResponse struct {
	Transactions []*model.Transaction `json:"transactions,omitempty"`
}

func (w *WalletServiceImpl) GetTransactions(ctx context.Context, req *GetTransactionsRequest) (*GetTransactionsResponse, error) {
	limit, offset := 10, 0 // default
	if req.Limit != nil {
		limit = *req.Limit
	}
	if req.Offset != nil {
		offset = *req.Offset
	}
	transactions, err := w.PostgreDal.GetTransactions(ctx, req.UserId, req.FromDate, req.ToDate, limit, offset)
	if err != nil {
		log.Printf("[WalletServiceImpl][GetTransactions] failed to get transactions, dbErr: %v", err)
		return nil, err
	}
	log.Printf("[WalletServiceImpl][GetTransactions] successfully getTransactions for user: %s, transactions: %v", req.UserId, transactions)
	return &GetTransactionsResponse{
		Transactions: transactions,
	}, nil
}
