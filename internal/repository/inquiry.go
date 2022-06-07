package repository

import (
	"context"
	"discount_service/internal/dto"
)

func (repo *Repo) GetDiscountsByID(ctx context.Context, walletID string) ([]dto.DiscountData, error) {
	walletDiscounts := []dto.DiscountData{}
	for k := range repo.DB {
		if walletID == repo.DB[k].WalletID {
			walletDiscounts = append(walletDiscounts, repo.DB[k])
		}
	}

	return walletDiscounts, nil
}

func (repo *Repo) GetDiscounts(ctx context.Context) (map[string]dto.DiscountData, int, int, error) {
	var usedDiscountCodes, totalDiscountCodes int
	for k := range repo.DB {
		totalDiscountCodes++
		if repo.DB[k].WalletID != "" {
			usedDiscountCodes++
		}
	}

	return repo.DB, usedDiscountCodes, totalDiscountCodes, nil
}
