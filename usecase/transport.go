package usecase

import (
	"context"
	"encoding/json"
	"net/http"

	"discount_service/usecase/dto"
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeChargeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.ChargeWalletRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
