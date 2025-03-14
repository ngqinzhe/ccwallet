package service

import "context"

type WithdrawRequest struct {
	UserId            int64   `json:"user_id"`
	Amount            float64 `json:"amount"`
	Currency          string  `json:"currency"`
	BankAccountNumber string  `json:"bank_account_number"`
	BankName          string  `json:"bank_name"`
}

type WithdrawResponse struct {
	Amount            float64 `json:"amount,omitempty"`
	Currency          string  `json:"currency,omitempty"`
	BankAccountNumber string  `json:"bank_account_number,omitempty"`
	BankName          string  `json:"bank_name,omitempty"`
	Error             error   `json:"error,omitempty"`
}

func (a *AccountService) Withdraw(ctx context.Context, req *WithdrawRequest) *WithdrawResponse {

}
