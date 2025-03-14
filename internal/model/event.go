package model

import (
	"encoding/json"
	"time"
)

// Event is the representation of transactional events that happen in this particular user
type EventType int

const (
	EventType_Deposit = iota + 1
	EventType_Withdraw
	EventType_Transfer
	EventType_GetAccountBalance
	EventType_GetTransactionHistory
)

type Event struct {
	Id        int64           `json:"id"`
	UserId    int64           `json:"user_id"`
	EventType EventType       `json:"event_type"`
	EventData json.RawMessage `json:"event_data"`
	CreatedAt time.Time       `json:"created_at"`
}

type DepositEvent struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type WithdrawEvent struct {
	Amount            float64 `json:"amount"`
	Currency          string  `json:"currency"`
	BankAccountNumber string  `json:"bank_account_number"`
	BankName          string  `json:"bank_name"`
}

type TransferEvent struct {
	FromUserId int64   `json:"from_user_id"`
	ToUserId   int64   `json:"to_user_id"`
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
}
