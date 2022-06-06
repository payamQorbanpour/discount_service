package dto

type (
	ChargeWalletResponse struct {
		ID      string `json:"id"`
		Balance int    `json:"balance"`
	}

	DiscountData struct {
		Code      string `json:"code"`
		WalletID  string `json:"wallet_id"`
		Amount    int    `json:"amount"`
		CreatedAt string
		UpdatedAt string
	}

	GetDiscountsByIDResponse struct {
		Discounts []DiscountData `json:"discounts"`
	}

	GetDiscountsResponse struct {
		Discounts map[string]DiscountData `json:"discounts"`
	}
)
