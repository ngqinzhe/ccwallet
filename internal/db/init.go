package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/ngqinzhe/ccwallet/internal/model"
)

type PostgreDal interface {
	Deposit(ctx context.Context, userId string, amount float64) error
	Withdraw(ctx context.Context, userId string, amount float64) error
	Transfer(ctx context.Context, fromUserId, toUserId string, amount float64) error
	GetWalletBalance(ctx context.Context, userId string) (float64, error)
	GetTransactions(ctx context.Context, userId string, from, to time.Time) ([]*model.Transaction, error)
	AddTransaction(ctx context.Context, userId, transactionType string, transactionData json.RawMessage) error
}

type PostgreDalImpl struct {
	client *sql.DB
}

func NewPostgreDal(ctx context.Context) PostgreDal {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	return &PostgreDalImpl{
		client: db,
	}
}
