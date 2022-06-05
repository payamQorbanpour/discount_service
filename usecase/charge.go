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
	ChargeWallet(ctx context.Context, wallet Wallet) error
}

type WebAPI interface {
	WalletChargeRequest(ctx context.Context, wallet Wallet) (res dto.ChargeWalletResponse, err error)
}
