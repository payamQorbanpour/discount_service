package dto

type (
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
