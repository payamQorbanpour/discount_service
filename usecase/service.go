package usecase

import "context"

type Service interface {
	ChargeWallet(ctx context.Context, id string, balance int) (totalBalance int, err error)
}
