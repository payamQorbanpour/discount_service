package usecase

import (
	"context"
	"discount_service/usecase/dto"
	"errors"
	"strconv"
	"time"

	"github.com/go-kit/kit/log"
)

type repo struct {
	db     map[string]dto.DiscountData
	logger log.Logger
}

func NewRepo(logger log.Logger) Repository {
	return &repo{
		db:     map[string]dto.DiscountData{},
		logger: logger,
	}
}

func (repo *repo) InitiateDiscounts(ctx context.Context, count, amount int) error {
	for i := 0; i < count; i++ {
		code := strconv.Itoa(i)
		discountData := dto.DiscountData{
			Code:      code,
			Amount:    amount,
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		}

		repo.db[code] = discountData
	}

	return nil
}

func (repo *repo) UpdateDiscount(ctx context.Context, code, walletID string) (*dto.DiscountData, error) {
	if !repo.checkDiscountExistance(ctx, code) {
		return nil, errors.New("discount code not found")
	}

	if repo.discountValidation(ctx, code) {
		return nil, errors.New("discount used")
	}

	discountData := repo.db[code]
	discountData.WalletID = walletID
	discountData.UpdatedAt = time.Now().Format(time.RFC3339)
	repo.db[code] = discountData
	return &discountData, nil
}

func (repo *repo) GetDiscountsByID(ctx context.Context, walletID string) ([]dto.DiscountData, error) {
	walletDiscounts := []dto.DiscountData{}
	for k := range repo.db {
		if walletID == repo.db[k].WalletID {
			walletDiscounts = append(walletDiscounts, repo.db[k])
		}
	}

	return walletDiscounts, nil
}

func (repo *repo) GetDiscounts(ctx context.Context) (map[string]dto.DiscountData, int, int, error) {
	var usedDiscountCodes, totalDiscountCodes int
	for k := range repo.db {
		totalDiscountCodes++
		if repo.db[k].WalletID != "" {
			usedDiscountCodes++
		}
	}

	return repo.db, usedDiscountCodes, totalDiscountCodes, nil
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
