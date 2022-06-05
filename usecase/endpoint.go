package usecase

import (
	"context"

	"discount_service/usecase/dto"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	ChargeWallet endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		ChargeWallet: makeChargeWalletEndpoint(s),
	}
}

func makeChargeWalletEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.ChargeWalletRequest)
		totalBalance, err := s.ChargeWallet(ctx, req.ID, req.Amount)
		return dto.ChargeWalletRequest{ID: req.ID, Amount: totalBalance}, err
	}
}
