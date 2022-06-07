package mockdata

import (
	"context"
	"discount_service/internal/dto"
	"discount_service/internal/webapi"

	"github.com/stretchr/testify/mock"
)

type MockWebAPI struct {
	mock.Mock
}

func (m *MockWebAPI) WalletChargeRequest(ctx context.Context, wallet webapi.Wallet) (dto.ChargeWalletResponse, error) {
	args := m.Called(ctx, wallet)
	return args.Get(0).(dto.ChargeWalletResponse), args.Error(1)
}
