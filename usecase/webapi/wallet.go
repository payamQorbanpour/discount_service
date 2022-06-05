package webapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/go-kit/kit/log"

	"discount_service/usecase"
	"discount_service/usecase/dto"
)

type web struct {
	logger log.Logger
}

func NewWebAPI(logger log.Logger) usecase.WebAPI {
	return &web{
		logger: logger,
	}
}

func (web *web) WalletChargeRequest(ctx context.Context, wallet usecase.Wallet) (resp dto.ChargeWalletResponse, err error) {
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
