package dto

type (
	DiscountsInitiateRequest struct {
		Count  int `json:"count"`
		Amount int `json:"amount"`
	}

	DiscountRequest struct {
		Code     string `json:"code"`
		WalletID string `json:"wallet_id"`
	}

	ChargeWalletRequest struct {
		ID     string `json:"id"`
		Amount int    `json:"amount"`
	}

	GetDiscountsByIDRequest struct {
		ID string `json:"id"`
	}

	GetDiscountsRequest struct {
	}
)
