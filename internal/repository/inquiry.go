package repository

import (
	"context"
	"discount_service/internal/dto"
)

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
