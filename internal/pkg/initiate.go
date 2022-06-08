package pkg

import (
	"context"
	"discount_service/internal/dto"

	"github.com/go-kit/log"
)

func (s DiscountService) InitiateDiscounts(ctx context.Context, count, amount int) (res *dto.GetDiscountsResponse, err error) {
	logger := log.With(s.logger, "method", "InitiateDiscounts")

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
