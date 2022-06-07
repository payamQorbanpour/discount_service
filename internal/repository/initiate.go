package repository

import (
	"context"
	"discount_service/internal/dto"
	"strconv"
	"time"
)

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
