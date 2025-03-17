package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/ngqinzhe/ccwallet/internal/model"
	"github.com/ngqinzhe/ccwallet/internal/util"
)

// mockgen -source=init.go -destination=../mocks/mock_dal.go -package=mocks
type PostgreDal interface {
	Deposit(ctx context.Context, userId string, amount float64) error
	Withdraw(ctx context.Context, userId string, amount float64) error
	Transfer(ctx context.Context, fromUserId, toUserId string, amount float64) error
	GetWalletBalance(ctx context.Context, userId string) (float64, error)
	GetTransactions(ctx context.Context, userId string, from, to time.Time, limit, offset int) ([]*model.Transaction, error)
}

type PostgreDalImpl struct {
	client *sql.DB
}

func NewPostgreDal(ctx context.Context, config *util.PostgreSqlCredentials) PostgreDal {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=localhost sslmode=disable",
		config.Username,
		config.Database,
		config.Password)
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
