package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"discount_service/internal/dto"
)

func DecodeChargeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.DiscountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
