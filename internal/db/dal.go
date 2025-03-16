package db

import (
	"context"
	"fmt"
	"time"

	"github.com/ngqinzhe/ccwallet/internal/model"
)

func (p *PostgreDalImpl) Deposit(ctx context.Context, userId string, amount float64) error {
	var err error
	tx, err := p.client.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func(err *error) {
		if err != nil {
			tx.Rollback()
		}
	}(&err)

	row := tx.QueryRowContext(ctx, "SELECT balance FROM wallet WHERE user_id = ? FOR UPDATE", userId)
	if row.Err() != nil {
		return row.Err()
	}
	var balance float64
	if err = row.Scan(&balance); err != nil {
		return err
	}
	balance += amount
	if _, err = tx.ExecContext(ctx, "UPDATE wallet SET balance = ? WHERE user_id = ?", balance, userId); err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, "INSERT INTO transaction (user_id, transaction_type, transaction_data) VALUES (?, ?, ?)",
		userId, model.TransactionType_Deposit, fmt.Sprintf("user: %s deposited %f to wallet", userId, amount)); err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgreDalImpl) Withdraw(ctx context.Context, userId string, amount float64) error {
	var err error

	tx, err := p.client.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func(err *error) {
		if err != nil {
			tx.Rollback()
		}
	}(&err)

	row := tx.QueryRowContext(ctx, "SELECT balance FROM wallet WHERE user_id = ? FOR UPDATE", userId)
	if row.Err() != nil {
		return row.Err()
	}
	var balance float64
	if err = row.Scan(&balance); err != nil {
		return err
	}
	balance -= amount
	if _, err = tx.ExecContext(ctx, "UPDATE wallet SET balance = ? WHERE user_id = ?", balance, userId); err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, "INSERT INTO transaction (user_id, transaction_type, transaction_data) VALUES (?, ?, ?)",
		userId, model.TransactionType_Withdraw, fmt.Sprintf("user: %s withdraw %f from wallet", userId, amount)); err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgreDalImpl) Transfer(ctx context.Context, fromUserId, toUserId string, amount float64) error {
	var err error

	tx, err := p.client.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func(err *error) {
		if err != nil {
			tx.Rollback()
		}
	}(&err)

	var fromUserBalance, toUserBalance float64

	row := tx.QueryRowContext(ctx, "SELECT balance FROM wallet WHERE user_id = ? FOR UPDATE", fromUserId)
	if row.Err() != nil {
		return row.Err()
	}
	if err = row.Scan(&fromUserBalance); err != nil {
		return err
	}

	row = tx.QueryRowContext(ctx, "SELECT balance FROM wallet WHERE user_id = ? FOR UPDATE", toUserId)
	if row.Err() != nil {
		return row.Err()
	}
	if err = row.Scan(&toUserBalance); err != nil {
		return err
	}

	fromUserBalance -= amount
	toUserBalance += amount

	if _, err = tx.ExecContext(ctx, "UPDATE wallet SET balance = ? WHERE user_id = ?", fromUserBalance, fromUserId); err != nil {
		return err
	}
	if _, err = tx.ExecContext(ctx, "UPDATE wallet SET balance = ? WHERE user_id = ?", toUserBalance, toUserId); err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, "INSERT INTO transaction (user_id, transaction_type, transaction_data) VALUES (?, ?, ?)",
		fromUserId, model.TransactionType_Withdraw, fmt.Sprintf("user: %s transfered %f from wallet to user: %s", fromUserId, amount, toUserId)); err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgreDalImpl) GetWalletBalance(ctx context.Context, userId string) (float64, error) {
	var balance float64
	row := p.client.QueryRowContext(ctx, "SELECT balance FROM wallet WHERE user_id = ?", userId)
	if row.Err() != nil {
		return 0, row.Err()
	}
	if err := row.Scan(&balance); err != nil {
		return 0, err
	}
	return balance, nil
}

func (p *PostgreDalImpl) GetTransactions(ctx context.Context, userId string, from, to time.Time) ([]*model.Transaction, error) {
	var transactions []*model.Transaction

	query := "SELECT * FROM transaction WHERE user_id = ?"
	queryParams := []interface{}{userId}
	if to.After(from) {
		query += " AND created_at >= ? AND created_at <= ?"
		queryParams = append(queryParams, from, to)
	}

	rows, err := p.client.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	if err := rows.Scan(&transactions); err != nil {
		return nil, err
	}
	return transactions, nil
}
