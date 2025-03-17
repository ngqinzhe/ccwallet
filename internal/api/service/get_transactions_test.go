package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ngqinzhe/ccwallet/internal/mocks"
	"github.com/ngqinzhe/ccwallet/internal/model"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetTransactions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mockDal := mocks.NewMockPostgreDal(ctrl)
	mockCache := mocks.NewMockRedisCache(ctrl)

	now := time.Now()

	t.Run("happy case", func(t *testing.T) {
		mockDal.EXPECT().GetTransactions(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*model.Transaction{
			{
				Id:          1,
				UserId:      "test_user",
				Type:        model.TransactionType_Deposit,
				Information: "deposited $1000",
				CreatedAt:   now,
			},
		}, nil).Times(1)

		w := NewWalletService(ctx, mockDal, mockCache)
		resp, err := w.GetTransactions(ctx, &GetTransactionsRequest{
			UserId:   "test_user",
			FromDate: now,
			ToDate:   now,
		})
		require.Nil(t, err)
		require.Equal(t, &GetTransactionsResponse{
			Transactions: []*model.Transaction{
				{
					Id:          1,
					UserId:      "test_user",
					Type:        model.TransactionType_Deposit,
					Information: "deposited $1000",
					CreatedAt:   now,
				},
			},
		}, resp)
	})

	t.Run("db fail", func(t *testing.T) {
		mockDal.EXPECT().GetTransactions(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("db fail")).Times(1)

		w := NewWalletService(ctx, mockDal, mockCache)
		resp, err := w.GetTransactions(ctx, &GetTransactionsRequest{
			UserId:   "test_user",
			FromDate: now,
			ToDate:   now,
		})
		require.NotNil(t, err)
		require.Nil(t, resp)
	})
}
