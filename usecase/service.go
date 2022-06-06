package usecase

import (
	"context"

	"discount_service/usecase/dto"
)

type Service interface {
	InitiateDiscounts(ctx context.Context, count, amount int) error
	ChargeWallet(ctx context.Context, code, walletID string) (res dto.ChargeWalletResponse, err error)
	GetDiscountsByID(ctx context.Context, id string) (res dto.GetDiscountsByIDResponse, err error)
	GetDiscounts(ctx context.Context) (res dto.GetDiscountsResponse, err error)
}
