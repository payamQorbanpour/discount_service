package endpoint

import (
	"context"
	"discount_service/internal/dto"
	"discount_service/internal/service"

	"github.com/go-kit/kit/endpoint"
)

func makeChargeWalletEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.DiscountRequest)
		res, err := s.ChargeWallet(ctx, req.Code, req.WalletID)
		return dto.ChargeWalletResponse{ID: res.ID, Balance: res.Balance}, err
	}
}
