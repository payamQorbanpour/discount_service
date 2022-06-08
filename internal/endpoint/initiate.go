package endpoint

import (
	"context"

	"discount_service/internal/dto"
	"discount_service/internal/pkg"

	"github.com/go-kit/kit/endpoint"
)

func makeInitiateDiscountsEndpoint(s pkg.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.DiscountsInitiateRequest)
		res, err := s.InitiateDiscounts(ctx, req.Count, req.Amount)
		if err != nil {
			return dto.GetDiscountsResponse{}, err
		}

		return dto.GetDiscountsResponse{Total: res.Total, Used: res.Used, Discounts: res.Discounts}, err
	}
}
