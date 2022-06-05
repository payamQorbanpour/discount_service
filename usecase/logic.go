package usecase

import (
	"context"

	"discount_service/usecase/dto"

	"github.com/go-kit/kit/log"
)

type service struct {
	repository Repository
	webAPI     WebAPI
	logger     log.Logger
}

func NewService(repo Repository, webAPI WebAPI, logger log.Logger) Service {
	return &service{
		repository: repo,
		webAPI:     webAPI,
		logger:     logger,
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

	s.repository.InsertDiscount(ctx, wallet.ID, wallet.Amount)

	logger.Log("Charge wallet", res.Balance)

	return res, nil
}
func (s service) GetDiscountsByID(ctx context.Context, id string) (res dto.GetDiscountsByIDResponse, err error) {
	logger := log.With(s.logger, "method", "GetDiscountsByID")
	dis, err := s.repository.GetDiscountsByID(ctx, id)
	if err != nil {
		logger.Log("err", err)
		return res, err
	}

	return dto.GetDiscountsByIDResponse{
		Discounts: dis,
	}, nil
}
