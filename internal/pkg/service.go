package pkg

import (
	"context"

	"discount_service/internal/dto"
	"discount_service/internal/repository"
	"discount_service/internal/webapi"

	"github.com/go-kit/log"
)

type DiscountService struct {
	repository repository.Repository
	webAPI     webapi.WebAPI
	logger     log.Logger
}

type Service interface {
	InitiateDiscounts(ctx context.Context, count, amount int) (res *dto.GetDiscountsResponse, err error)
	ChargeWallet(ctx context.Context, code, walletID string) (res *dto.ChargeWalletResponse, err error)
	GetDiscountsByID(ctx context.Context, id string) (res dto.GetDiscountsByIDResponse, err error)
	GetDiscounts(ctx context.Context) (res dto.GetDiscountsResponse, err error)
}

func NewService(repo repository.Repository, webAPI webapi.WebAPI, logger log.Logger) Service {
	return &DiscountService{
		repository: repo,
		webAPI:     webAPI,
		logger:     logger,
	}
}
