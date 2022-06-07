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

func (s service) InitiateDiscounts(ctx context.Context, count, amount int) error {
	logger := log.With(s.logger, "method", "InitiateDiscounts")

	err := s.repository.InitiateDiscounts(ctx, count, amount)
	if err != nil {
		logger.Log("err", err)
		return err
	}

	return nil
}

func (s service) ChargeWallet(ctx context.Context, code, walletID string) (res dto.ChargeWalletResponse, err error) {
	logger := log.With(s.logger, "method", "ChargeWallet")

	discountData, err := s.repository.UpdateDiscount(ctx, code, walletID)
	if err != nil {
		logger.Log("err", err)
		return res, err
	}

	wallet := Wallet{
		ID:     walletID,
		Amount: discountData.Amount,
	}

	res, err = s.webAPI.WalletChargeRequest(ctx, wallet)
	if err != nil {
		logger.Log("err", err)
		return res, err
	}

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

func (s service) GetDiscounts(ctx context.Context) (res dto.GetDiscountsResponse, err error) {
	logger := log.With(s.logger, "method", "GetDiscounts")
	dis, usedDiscountCodes, totalDiscountCodes, err := s.repository.GetDiscounts(ctx)
	if err != nil {
		logger.Log("err", err)
		return res, err
	}

	return dto.GetDiscountsResponse{
		Total:     totalDiscountCodes,
		Used:      usedDiscountCodes,
		Discounts: dis,
	}, nil
}
