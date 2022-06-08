package endpoint

import (
	"discount_service/internal/pkg"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	InitiateDiscounts endpoint.Endpoint
	ChargeWallet      endpoint.Endpoint
	GetDiscountsByID  endpoint.Endpoint
	GetDiscounts      endpoint.Endpoint
}

func MakeEndpoints(s pkg.Service) Endpoints {
	return Endpoints{
		InitiateDiscounts: makeInitiateDiscountsEndpoint(s),
		ChargeWallet:      makeChargeWalletEndpoint(s),
		GetDiscountsByID:  makeGetDiscountsByID(s),
		GetDiscounts:      makeGetDiscounts(s),
	}
}
