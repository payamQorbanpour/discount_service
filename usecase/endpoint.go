package usecase

import (
	"context"

	"discount_service/usecase/dto"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	ChargeWallet     endpoint.Endpoint
	GetDiscountsByID endpoint.Endpoint
	GetDiscounts     endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		ChargeWallet:     makeChargeWalletEndpoint(s),
		GetDiscountsByID: makeGetDiscountsByID(s),
		GetDiscounts:     makeGetDiscounts(s),
	}
}

func makeChargeWalletEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.ChargeWalletRequest)
		res, err := s.ChargeWallet(ctx, req.ID, req.Amount)
		return dto.ChargeWalletResponse{ID: res.ID, Balance: res.Balance}, err
	}
}

func makeGetDiscountsByID(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.GetDiscountsByIDRequest)
		res, err := s.GetDiscountsByID(ctx, req.ID)
		return dto.GetDiscountsByIDResponse{Discounts: res.Discounts}, err
	}
}

func makeGetDiscounts(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(dto.GetDiscountsRequest)
		res, err := s.GetDiscounts(ctx)
		return dto.GetDiscountsResponse{Discounts: res.Discounts}, err
	}
}
