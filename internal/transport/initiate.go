package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"discount_service/internal/dto"
)

func DecodeInitiateDiscounts(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.DiscountsInitiateRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}
