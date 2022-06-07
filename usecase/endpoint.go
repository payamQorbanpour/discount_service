package usecase

import (
	"context"

	"discount_service/usecase/dto"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	InitiateDiscounts endpoint.Endpoint
	ChargeWallet      endpoint.Endpoint
	GetDiscountsByID  endpoint.Endpoint
	GetDiscounts      endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		InitiateDiscounts: makeInitiateDiscountsEndpoint(s),
		ChargeWallet:      makeChargeWalletEndpoint(s),
		GetDiscountsByID:  makeGetDiscountsByID(s),
		GetDiscounts:      makeGetDiscounts(s),
	}
}

func makeInitiateDiscountsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.DiscountsInitiateRequest)
		err = s.InitiateDiscounts(ctx, req.Count, req.Amount)
		return dto.ChargeWalletResponse{}, err
	}
}

func makeChargeWalletEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.DiscountRequest)
		res, err := s.ChargeWallet(ctx, req.Code, req.WalletID)
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
		return dto.GetDiscountsResponse{Total: res.Total, Used: res.Used, Discounts: res.Discounts}, err
	}
}
