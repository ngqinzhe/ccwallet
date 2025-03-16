package model

import (
	"time"
)

// Transaction is the representation of transactional Transactions that happen in this particular user
type TransactionType int

const (
	TransactionType_Deposit = iota + 1
	TransactionType_Withdraw
	TransactionType_Transfer
)

type Transaction struct {
	Id          int64           `json:"id"`
	UserId      string          `json:"user_id"`
	Type        TransactionType `json:"type"`
	Information string          `json:"information"`
	CreatedAt   time.Time       `json:"created_at"`
}
