package service

import (
	"context"
	"errors"
	"testing"

	"github.com/ngqinzhe/ccwallet/internal/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestDeposit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mockDal := mocks.NewMockPostgreDal(ctrl)
	mockCache := mocks.NewMockRedisCache(ctrl)

	t.Run("happy case", func(t *testing.T) {
		mockDal.EXPECT().Deposit(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)

		w := NewWalletService(ctx, mockDal, mockCache)
		resp, err := w.Deposit(ctx, &DepositRequest{
			UserId: "test_user",
			Amount: 1000,
		})
		require.Nil(t, err)
		require.Equal(t, &DepositResponse{
			Amount: 1000,
		}, resp)
	})

	t.Run("dal failed", func(t *testing.T) {
		mockDal.EXPECT().Deposit(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("failed")).Times(1)
		w := NewWalletService(ctx, mockDal, mockCache)
		resp, err := w.Deposit(ctx, &DepositRequest{
			UserId: "test_user",
			Amount: 1000,
		})
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

}
