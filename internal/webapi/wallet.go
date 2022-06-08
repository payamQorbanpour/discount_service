package webapi

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"discount_service/internal/dto"
	"discount_service/internal/model"

	"github.com/go-kit/log"
)

type web struct {
	logger log.Logger
}

type Wallet struct {
	ID        string `json:"id,omitempty"`
	Amount    int    `json:"amount"`
	CreatedAt time.Time
}

func NewWebAPI(logger log.Logger) WebAPI {
	return &web{
		logger: logger,
	}
}

func (web *web) WalletChargeRequest(ctx context.Context, wallet Wallet) (resp *dto.ChargeWalletResponse, err error) {
	logger := log.With(web.logger, "method", "WalletChargeRequest")

	url := "http://localhost:8080/charge"

	payload, _ := json.Marshal(&dto.ChargeWalletRequest{
		ID:     wallet.ID,
		Amount: wallet.Amount,
	})

	res, err := http.Post(url, "Content-Type:application/json", bytes.NewBuffer(payload))
	if err != nil {
		logger.Log("err", err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Log("err", err)
		return nil, err
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		logger.Log("err", body)
		return nil, model.ErrServiceUnavailable
	}

	return resp, nil
}
