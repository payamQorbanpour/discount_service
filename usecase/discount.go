package usecase

import (
	"context"
	"time"

	"discount_service/usecase/dto"
)

type Wallet struct {
	ID        string `json:"id,omitempty"`
	Amount    int    `json:"amount"`
	CreatedAt time.Time
}

type Repository interface {
	InitiateDiscounts(ctx context.Context, count, amount int) error
	UpdateDiscount(ctx context.Context, code, walletID string) (*dto.DiscountData, error)
	GetDiscountsByID(ctx context.Context, walletID string) ([]dto.DiscountData, error)
	GetDiscounts(ctx context.Context) (map[string]dto.DiscountData, int, int, error)
}

type WebAPI interface {
	WalletChargeRequest(ctx context.Context, wallet Wallet) (res dto.ChargeWalletResponse, err error)
}
