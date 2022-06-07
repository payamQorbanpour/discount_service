package webapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"discount_service/internal/dto"

	"github.com/go-kit/kit/log"
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

func (web *web) WalletChargeRequest(ctx context.Context, wallet Wallet) (resp dto.ChargeWalletResponse, err error) {
	logger := log.With(web.logger, "method", "WalletChargeRequest")

	url := "http://localhost:8085/charge"

	payload, _ := json.Marshal(&dto.ChargeWalletRequest{
		ID:     wallet.ID,
		Amount: wallet.Amount,
	})

	res, err := http.Post(url, "Content-Type:application/json", bytes.NewBuffer(payload))
	if err != nil {
		logger.Log("err", err)
		return resp, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Log("err", err)
		return resp, err
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		logger.Log("err", body)
		return resp, errors.New(string(body))
	}

	return resp, nil
}
