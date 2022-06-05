package usecase

import (
	"context"

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

func (s service) ChargeWallet(ctx context.Context, id string, amount int) (totalBalance int, err error) {
	logger := log.With(s.logger, "method", "ChargeWallet")

	wallet := Wallet{
		ID:     id,
		Amount: amount,
	}

	totalBalance, err = s.webAPI.WalletChargeRequest(ctx, wallet)
	if err != nil {
		logger.Log("err", err)
		return -1, err
	}

	logger.Log("Charge wallet", totalBalance)

	return totalBalance, nil
}
