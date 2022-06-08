package webapi

import (
	"context"

	"discount_service/internal/dto"
)

type WebAPI interface {
	WalletChargeRequest(ctx context.Context, wallet Wallet) (res *dto.ChargeWalletResponse, err error)
}
