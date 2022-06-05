package usecase

import (
	"context"

	"discount_service/usecase/dto"
)

type Service interface {
	ChargeWallet(ctx context.Context, id string, balance int) (res dto.ChargeWalletResponse, err error)
}
