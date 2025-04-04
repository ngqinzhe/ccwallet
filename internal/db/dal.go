package db

import (
	"context"
	"errors"
	"fmt"
	"log"
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

	row := tx.QueryRowContext(ctx, "SELECT balance FROM public.wallet WHERE user_id = $1 FOR UPDATE", userId)
	if row.Err() != nil {
		return row.Err()
	}
	var balance float64
	if err = row.Scan(&balance); err != nil {
		return err
	}
	balance += amount
	if _, err = tx.ExecContext(ctx, "UPDATE wallet SET balance = $1 WHERE user_id = $2", balance, userId); err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, "INSERT INTO transaction (user_id, type, information) VALUES ($1, $2, $3)",
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

	row := tx.QueryRowContext(ctx, "SELECT balance FROM public.wallet WHERE user_id = $1 FOR UPDATE", userId)
	if row.Err() != nil {
		return row.Err()
	}
	var balance float64
	if err = row.Scan(&balance); err != nil {
		return err
	}
	if balance < amount {
		return errors.New("user does not have sufficient balance to withdraw")
	}

	balance -= amount
	if _, err = tx.ExecContext(ctx, "UPDATE wallet SET balance = $1 WHERE user_id = $2", balance, userId); err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, "INSERT INTO transaction (user_id, type, information) VALUES ($1, $2, $3)",
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

	row := tx.QueryRowContext(ctx, "SELECT balance FROM public.wallet WHERE user_id = $1 FOR UPDATE", fromUserId)
	if row.Err() != nil {
		return row.Err()
	}
	if err = row.Scan(&fromUserBalance); err != nil {
		return err
	}
	if fromUserBalance < amount {
		return errors.New("user has insufficient amount")
	}

	row = tx.QueryRowContext(ctx, "SELECT balance FROM public.wallet WHERE user_id = $1 FOR UPDATE", toUserId)
	if row.Err() != nil {
		return row.Err()
	}
	if err = row.Scan(&toUserBalance); err != nil {
		return err
	}

	fromUserBalance -= amount
	toUserBalance += amount

	if _, err = tx.ExecContext(ctx, "UPDATE wallet SET balance = $1 WHERE user_id = $2", fromUserBalance, fromUserId); err != nil {
		return err
	}
	if _, err = tx.ExecContext(ctx, "UPDATE wallet SET balance = $1 WHERE user_id = $2", toUserBalance, toUserId); err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, "INSERT INTO transaction (user_id, type, information) VALUES (?, ?, ?)",
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
	row := p.client.QueryRowContext(ctx, "SELECT balance FROM public.wallet WHERE user_id = $1", userId)
	if row.Err() != nil {
		log.Print("fail1")
		return 0, row.Err()
	}
	if err := row.Scan(&balance); err != nil {
		log.Print("fail2")
		return 0, err
	}
	return balance, nil
}

func (p *PostgreDalImpl) GetTransactions(ctx context.Context, userId string, from, to time.Time, limit, offset int) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	query := "SELECT * FROM transaction WHERE user_id = $1"
	queryParams := []interface{}{userId}

	argIndex := 2
	if to.After(from) {
		query += fmt.Sprintf(" AND created_at >= $%d AND created_at <= $%d", argIndex, argIndex+1)
		queryParams = append(queryParams, from, to)
		argIndex += 2
	}
	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	rows, err := p.client.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		transaction := &model.Transaction{}
		if err := rows.Scan(&transaction.Id, &transaction.UserId, &transaction.Type, &transaction.Information, &transaction.CreatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
