package pkg

import (
	"context"
	"discount_service/internal/dto"
	"discount_service/internal/model"

	"github.com/go-kit/log"
)

func (s DiscountService) InitiateDiscounts(ctx context.Context, count, amount int) (res *dto.GetDiscountsResponse, err error) {
	logger := log.With(s.logger, "method", "InitiateDiscounts")

	if s.amountValidation(ctx, amount) {
		return nil, model.ErrNotValidAmount
	}

	err = s.repository.InitiateDiscounts(ctx, count, amount)
	if err != nil {
		logger.Log("err", err)
		return nil, err
	}

	discounts, err := s.GetDiscounts(ctx)
	if err != nil {
		logger.Log("err", err)
		return nil, err
	}

	res = &dto.GetDiscountsResponse{
		Total:     discounts.Total,
		Used:      discounts.Used,
		Discounts: discounts.Discounts,
	}

	return res, nil
}

func (s DiscountService) amountValidation(ctx context.Context, amount int) bool {
	return amount <= 0
}
