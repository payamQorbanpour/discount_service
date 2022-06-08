package repository

import (
	"context"

	"discount_service/internal/dto"

	"github.com/go-kit/log"
)

type Repo struct {
	DB     map[string]dto.DiscountData
	logger log.Logger
}

type Repository interface {
	InitiateDiscounts(ctx context.Context, count, amount int) error
	UpdateDiscount(ctx context.Context, code, walletID string) (*dto.DiscountData, error)
	UpdateDiscountRollback(ctx context.Context, code, walletID string) (*dto.DiscountData, error)
	GetDiscountsByID(ctx context.Context, walletID string) ([]dto.DiscountData, error)
	GetDiscounts(ctx context.Context) (map[string]dto.DiscountData, int, int, error)
}

func NewRepo(logger log.Logger) Repository {
	return &Repo{
		DB:     map[string]dto.DiscountData{},
		logger: logger,
	}
}

func (repo *Repo) checkDiscountExistance(ctx context.Context, code string) bool {
	if _, exists := repo.DB[code]; exists {
		return true
	}

	return false
}

func (repo *Repo) discountValidation(ctx context.Context, code string) bool {
	return repo.DB[code].WalletID != ""
}
