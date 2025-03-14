package service

import "context"

type WithdrawRequest struct {
	UserId            string  `json:"user_id"`
	Amount            float64 `json:"amount"`
	BankAccountNumber string  `json:"bank_account_number"`
	BankName          string  `json:"bank_name"`
}

type WithdrawResponse struct {
	Amount float64 `json:"amount,omitempty"`

	BankAccountNumber string `json:"bank_account_number,omitempty"`
	BankName          string `json:"bank_name,omitempty"`
	Error             error  `json:"error,omitempty"`
}

func (a *AccountServiceImpl) Withdraw(ctx context.Context, req *WithdrawRequest) (*WithdrawResponse, error) {
}
