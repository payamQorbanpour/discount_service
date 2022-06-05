package usecase

import (
	"context"

	"discount_service/usecase/dto"

	"github.com/go-kit/kit/log"
)

type service struct {
	webAPI WebAPI
	logger log.Logger
}

func NewService(webAPI WebAPI, logger log.Logger) Service {
	return &service{
		webAPI: webAPI,
		logger: logger,
	}
}

func (s service) ChargeWallet(ctx context.Context, id string, amount int) (res dto.ChargeWalletResponse, err error) {
	logger := log.With(s.logger, "method", "ChargeWallet")

	wallet := Wallet{
		ID:     id,
		Amount: amount,
	}

	res, err = s.webAPI.WalletChargeRequest(ctx, wallet)
	if err != nil {
		logger.Log("err", err)
		return res, err
	}

	logger.Log("Charge wallet", res.Balance)

	return res, nil
}
