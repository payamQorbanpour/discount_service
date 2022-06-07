package pkg

import (
	"context"
	"discount_service/internal/dto"
	"discount_service/internal/webapi"

	"github.com/go-kit/kit/log"
)

func (s *DiscountService) ChargeWallet(ctx context.Context, code, walletID string) (res dto.ChargeWalletResponse, err error) {
	logger := log.With(s.logger, "method", "ChargeWallet")

	discountData, err := s.repository.UpdateDiscount(ctx, code, walletID)
	if err != nil {
		logger.Log("err", err)
		return res, err
	}

	wallet := webapi.Wallet{
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
