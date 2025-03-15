package model

import (
	"encoding/json"
	"time"
)

// Transaction is the representation of transactional Transactions that happen in this particular user
type TransactionType int

const (
	TransactionType_Deposit = iota + 1
	TransactionType_Withdraw
	TransactionType_Transfer
	TransactionType_GetAccountBalance
	TransactionType_GetTransactionHistory
)

type Transaction struct {
	Id              int64           `json:"id"`
	UserId          string          `json:"user_id"`
	TransactionType TransactionType `json:"transaction_type"`
	TransactionData json.RawMessage `json:"transaction_data"`
	CreatedAt       time.Time       `json:"created_at"`
}

type DepositTransaction struct {
	Amount float64 `json:"amount"`
}

type WithdrawTransaction struct {
	Amount            float64 `json:"amount"`
	BankAccountNumber string  `json:"bank_account_number"`
	BankName          string  `json:"bank_name"`
}

type TransferTransaction struct {
	FromUserId int64   `json:"from_user_id"`
	ToUserId   int64   `json:"to_user_id"`
	Amount     float64 `json:"amount"`
}
