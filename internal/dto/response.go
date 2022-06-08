package dto

type (
	Error struct {
		Message   string `json:"message"`
		ErrorCode int    `json:"error_code"`
	}

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
		Total     int                     `json:"total"`
		Used      int                     `json:"used"`
		Discounts map[string]DiscountData `json:"discounts"`
	}
)
