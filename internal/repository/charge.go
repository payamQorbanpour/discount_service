package repository

import (
	"context"
	"time"

	"discount_service/internal/dto"
	"discount_service/internal/model"
)

func (repo *Repo) UpdateDiscount(ctx context.Context, code, walletID string) (*dto.DiscountData, error) {
	if !repo.checkDiscountExistance(ctx, code) {
		return nil, model.ErrDiscountNotFound
	}

	if repo.discountValidation(ctx, code) {
		return nil, model.ErrDiscountUsed
	}

	discountData := repo.DB[code]
	discountData.WalletID = walletID
	discountData.UpdatedAt = time.Now().Format(time.RFC3339)
	repo.DB[code] = discountData
	return &discountData, nil
}
