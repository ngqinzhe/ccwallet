package service

import (
	"context"
	"errors"
	"testing"

	"github.com/ngqinzhe/ccwallet/internal/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestWithdraw(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mockDal := mocks.NewMockPostgreDal(ctrl)
	mockCache := mocks.NewMockRedisCache(ctrl)

	t.Run("happy case", func(t *testing.T) {
		mockDal.EXPECT().Withdraw(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)

		w := NewWalletService(ctx, mockDal, mockCache)
		resp, err := w.Withdraw(ctx, &WithdrawRequest{
			UserId:            "test_user",
			Amount:            1000,
			BankAccountNumber: "12345",
			BankName:          "DBS",
		})
		require.Nil(t, err)
		require.Equal(t, &WithdrawResponse{
			Amount:            1000,
			BankAccountNumber: "12345",
			BankName:          "DBS",
		}, resp)
	})

	t.Run("withdraw failed", func(t *testing.T) {
		mockDal.EXPECT().Withdraw(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("withdraw failed")).Times(1)

		w := NewWalletService(ctx, mockDal, mockCache)
		resp, err := w.Withdraw(ctx, &WithdrawRequest{
			UserId:            "test_user",
			Amount:            1000,
			BankAccountNumber: "12345",
			BankName:          "DBS",
		})
		require.NotNil(t, err)
		require.Nil(t, resp)
	})
}
