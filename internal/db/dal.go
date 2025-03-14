package db

import (
	"context"
)

func (p *PostgreDalImpl) Deposit(ctx context.Context, userId string, amount float64) error {
	p.client.BeginTx(ctx, nil)
	_, err := p.client.Exec("UPDATE wallet SET balance = balance + ?", amount)
	if err != nil {
		return err
	}
	return nil
}
