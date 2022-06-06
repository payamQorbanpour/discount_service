package usecase

import (
	"context"
	"encoding/json"
	"net/http"

	"discount_service/usecase/dto"

	"github.com/gorilla/mux"
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeInitiateDiscounts(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.DiscountsInitiateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeChargeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.DiscountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeGetDiscountsByIDRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetDiscountsByIDRequest
	vars := mux.Vars(r)

	req = dto.GetDiscountsByIDRequest{
		ID: vars["id"],
	}

	return req, nil
}

func decodeGetDiscountsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetDiscountsRequest

	return req, nil
}
