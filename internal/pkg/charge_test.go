package pkg

import (
	"context"
	"os"
	"testing"
	"time"

	"discount_service/internal/dto"
	"discount_service/internal/repository"
	"discount_service/internal/webapi"
	"discount_service/mockdata"

	"github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var walletChargeRequestTests = struct {
	name       string
	repo       repository.Repo
	entry      *webapi.Wallet
	mockWebAPI dto.ChargeWalletResponse
	mockError  error
	want       dto.ChargeWalletResponse
	wantError  bool
}{
	name: "success",
	repo: repository.Repo{
		DB: map[string]dto.DiscountData{
			"1": {
				Code:      "1",
				WalletID:  "",
				Amount:    1000000,
				CreatedAt: time.Now().Format(time.RFC3339),
				UpdatedAt: time.Now().Format(time.RFC3339),
			},
		},
	},
	entry: nil,
	want: dto.ChargeWalletResponse{
		ID:      "",
		Balance: 0,
	},
	wantError: true,
}

func TestChargeWallet(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "discount test",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	mockWebAPI := new(mockdata.MockWebAPI)

	srv := &DiscountService{
		repository: &walletChargeRequestTests.repo,
		webAPI:     mockWebAPI,
		logger:     logger,
	}

	t.Run(walletChargeRequestTests.name, func(t *testing.T) {
		mockWebAPI.On("WalletChargeRequest", context.Background(), mock.Anything).Return(walletChargeRequestTests.mockWebAPI, walletChargeRequestTests.mockError)
		got, err := srv.ChargeWallet(context.Background(), "1", "09123456789")
		if err != nil {
			assert.Error(t, err)
		}
		assert.Equal(t, walletChargeRequestTests.want, got)
	})

	mockWebAPI.AssertExpectations(t)
}
