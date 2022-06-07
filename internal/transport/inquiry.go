package transport

import (
	"context"
	"net/http"

	"discount_service/internal/dto"

	"github.com/gorilla/mux"
)

func DecodeGetDiscountsByIDRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetDiscountsByIDRequest
	vars := mux.Vars(r)

	req = dto.GetDiscountsByIDRequest{
		ID: vars["id"],
	}

	return req, nil
}

func DecodeGetDiscountsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetDiscountsRequest

	return req, nil
}
