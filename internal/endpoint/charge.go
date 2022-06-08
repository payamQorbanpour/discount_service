package endpoint

import (
	"context"

	"discount_service/internal/dto"
	"discount_service/internal/pkg"

	"github.com/go-kit/kit/endpoint"
)

// @Summary Charge wallet
// @ID charge
// @Description Send charge wallet request with given amount and store using discount code
// @Accept json
// @Produce json
// @Tags Charge
// @Param request body dto.DiscountRequest true "Charge request"
// @Success 200 {object} dto.ChargeWalletResponse
// @Failure 404 {object} dto.Error
// @Router / [post]
// .
func makeChargeWalletEndpoint(s pkg.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.DiscountRequest)
		res, err := s.ChargeWallet(ctx, req.Code, req.WalletID)
		if err != nil {
			return dto.ChargeWalletResponse{}, err
		}

		return dto.ChargeWalletResponse{ID: res.ID, Balance: res.Balance}, err
	}
}
