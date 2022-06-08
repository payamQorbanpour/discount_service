package endpoint

import (
	"context"

	"discount_service/internal/dto"
	"discount_service/internal/pkg"

	"github.com/go-kit/kit/endpoint"
)

func makeGetDiscountsByID(s pkg.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.GetDiscountsByIDRequest)
		res, err := s.GetDiscountsByID(ctx, req.ID)
		return dto.GetDiscountsByIDResponse{Discounts: res.Discounts}, err
	}
}

func makeGetDiscounts(s pkg.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(dto.GetDiscountsRequest)
		res, err := s.GetDiscounts(ctx)
		return dto.GetDiscountsResponse{Total: res.Total, Used: res.Used, Discounts: res.Discounts}, err
	}
}
