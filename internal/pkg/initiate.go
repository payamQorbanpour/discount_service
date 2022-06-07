package pkg

import (
	"context"

	"github.com/go-kit/kit/log"
)

func (s DiscountService) InitiateDiscounts(ctx context.Context, count, amount int) error {
	logger := log.With(s.logger, "method", "InitiateDiscounts")

	err := s.repository.InitiateDiscounts(ctx, count, amount)
	if err != nil {
		logger.Log("err", err)
		return err
	}

	return nil
}
