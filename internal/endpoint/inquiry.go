package endpoint

import (
	"context"

	"discount_service/internal/dto"
	"discount_service/internal/pkg"

	"github.com/go-kit/kit/endpoint"
)

// @Summary Get discounts by id
// @ID getdiscountsbyid
// @Description Get all discounts a wallet get
// @Accept json
// @Produce json
// @Tags GetDiscounts
// @Param request body dto.GetDiscountsByIDRequest true "Get discount by id request"
// @Success 200 {object} dto.GetDiscountsByIDResponse
// @Failure 404 {object} dto.Error
// @Router /{id} [get]
// .
func makeGetDiscountsByID(s pkg.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.GetDiscountsByIDRequest)
		res, err := s.GetDiscountsByID(ctx, req.ID)
		return dto.GetDiscountsByIDResponse{Discounts: res.Discounts}, err
	}
}

// @Summary Get discounts
// @ID getdiscounts
// @Description Get all discounts
// @Accept json
// @Produce json
// @Tags GetDiscounts
// @Param request body dto.GetDiscountsRequest true "Get discount request"
// @Success 200 {object} dto.GetDiscountsResponse
// @Failure 404 {object} dto.Error
// @Router /{id} [get]
// .
func makeGetDiscounts(s pkg.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(dto.GetDiscountsRequest)
		res, err := s.GetDiscounts(ctx)
		return dto.GetDiscountsResponse{Total: res.Total, Used: res.Used, Discounts: res.Discounts}, err
	}
}
