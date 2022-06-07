package pkg

import (
	"context"
	"discount_service/internal/dto"

	"github.com/go-kit/kit/log"
)

func (s discountService) GetDiscountsByID(ctx context.Context, id string) (res dto.GetDiscountsByIDResponse, err error) {
	logger := log.With(s.logger, "method", "GetDiscountsByID")
	dis, err := s.repository.GetDiscountsByID(ctx, id)
	if err != nil {
		logger.Log("err", err)
		return res, err
	}

	return dto.GetDiscountsByIDResponse{
		Discounts: dis,
	}, nil
}

func (s discountService) GetDiscounts(ctx context.Context) (res dto.GetDiscountsResponse, err error) {
	logger := log.With(s.logger, "method", "GetDiscounts")
	dis, usedDiscountCodes, totalDiscountCodes, err := s.repository.GetDiscounts(ctx)
	if err != nil {
		logger.Log("err", err)
		return res, err
	}

	return dto.GetDiscountsResponse{
		Total:     totalDiscountCodes,
		Used:      usedDiscountCodes,
		Discounts: dis,
	}, nil
}
