package service

import (
	"context"
	"errors"
	"testing"

	"github.com/ngqinzhe/ccwallet/internal/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestTransfer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mockDal := mocks.NewMockPostgreDal(ctrl)
	mockCache := mocks.NewMockRedisCache(ctrl)

	t.Run("happy case", func(t *testing.T) {
		mockDal.EXPECT().Transfer(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)

		w := NewWalletService(ctx, mockDal, mockCache)
		resp, err := w.Transfer(ctx, &TransferRequest{
			FromUserId: "test_user",
			ToUserId:   "test_user2",
			Amount:     1000,
		})
		require.Nil(t, err)
		require.Equal(t, &TransferResponse{
			Amount: 1000,
		}, resp)
	})

	t.Run("withdraw failed", func(t *testing.T) {
		mockDal.EXPECT().Transfer(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("transfer failed")).Times(1)

		w := NewWalletService(ctx, mockDal, mockCache)
		resp, err := w.Transfer(ctx, &TransferRequest{
			FromUserId: "test_user",
			ToUserId:   "test_user2",
			Amount:     1000,
		})
		require.NotNil(t, err)
		require.Nil(t, resp)
	})
}
