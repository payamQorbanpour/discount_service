package endpoint

import (
	"context"
	"discount_service/internal/dto"
	"discount_service/internal/service"

	"github.com/go-kit/kit/endpoint"
)

func makeInitiateDiscountsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.DiscountsInitiateRequest)
		err = s.InitiateDiscounts(ctx, req.Count, req.Amount)
		return dto.ChargeWalletResponse{}, err
	}
}
