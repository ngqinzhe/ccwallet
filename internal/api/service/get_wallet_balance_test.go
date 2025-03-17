package service

import (
	"context"
	"errors"
	"testing"

	"github.com/ngqinzhe/ccwallet/internal/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetWalletBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mockDal := mocks.NewMockPostgreDal(ctrl)
	mockCache := mocks.NewMockRedisCache(ctrl)

	t.Run("happy case from cache", func(t *testing.T) {
		mockCache.EXPECT().Get(gomock.Any(), gomock.Any()).Return(float64(1000), nil).Times(1)

		w := NewWalletService(ctx, mockDal, mockCache)
		resp, err := w.GetWalletBalance(ctx, &GetWalletBalanceRequest{
			UserId: "test_user",
		})
		require.Nil(t, err)
		require.Equal(t, &GetWalletBalanceResponse{
			Balance: 1000,
		}, resp)
	})

	t.Run("fail from cache, read from db", func(t *testing.T) {
		mockCache.EXPECT().Get(gomock.Any(), gomock.Any()).Return("", errors.New("")).Times(1)
		mockDal.EXPECT().GetWalletBalance(gomock.Any(), gomock.Any()).Return(float64(1000), nil).Times(1)
		mockCache.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)

		w := NewWalletService(ctx, mockDal, mockCache)
		resp, err := w.GetWalletBalance(ctx, &GetWalletBalanceRequest{
			UserId: "test_user",
		})
		require.Nil(t, err)
		require.Equal(t, &GetWalletBalanceResponse{
			Balance: 1000,
		}, resp)
	})
}
