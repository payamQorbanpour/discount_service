package repository

import (
	"context"
	"discount_service/internal/dto"

	"github.com/go-kit/kit/log"
)

type repo struct {
	db     map[string]dto.DiscountData
	logger log.Logger
}

type Repository interface {
	InitiateDiscounts(ctx context.Context, count, amount int) error
	UpdateDiscount(ctx context.Context, code, walletID string) (*dto.DiscountData, error)
	GetDiscountsByID(ctx context.Context, walletID string) ([]dto.DiscountData, error)
	GetDiscounts(ctx context.Context) (map[string]dto.DiscountData, int, int, error)
}

func NewRepo(logger log.Logger) Repository {
	return &repo{
		db:     map[string]dto.DiscountData{},
		logger: logger,
	}
}

func (repo *repo) checkDiscountExistance(ctx context.Context, code string) bool {
	if _, exists := repo.db[code]; exists {
		return true
	}

	return false
}

func (repo *repo) discountValidation(ctx context.Context, code string) bool {
	return repo.db[code].WalletID != ""
}
