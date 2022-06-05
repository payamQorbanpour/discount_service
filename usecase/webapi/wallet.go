package webapi

import (
	"bytes"
	"context"
	"encoding/json"
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

func (web *web) WalletChargeRequest(ctx context.Context, wallet usecase.Wallet) (balance int, err error) {
	logger := log.With(web.logger, "method", "WalletChargeRequest")

	url := "http://localhost:8085/charge"

	payload, _ := json.Marshal(&dto.ChargeWalletRequest{
		ID:     wallet.ID,
		Amount: wallet.Amount,
	})

	resp, err := http.Post(url, "Content-Type:application/json", bytes.NewBuffer(payload))
	if err != nil {
		logger.Log("err", err)
		return -1, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Log("err", err)
		return -1, err
	}

	inquiry := &dto.ChargeWalletResponse{}
	if err := json.Unmarshal(body, inquiry); err != nil {
		logger.Log("err", err)
		return -1, err
	}

	return inquiry.Balance, nil
}
