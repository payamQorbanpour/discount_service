package repository

import (
	"context"
	"discount_service/internal/dto"
	"errors"
	"time"
)

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
