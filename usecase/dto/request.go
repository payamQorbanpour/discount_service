package dto

type (
	ChargeWalletRequest struct {
		ID     string `json:"id"`
		Amount int    `json:"amount"`
	}
)

type (
	GetDiscountsByIDRequest struct {
		ID string `json:"id"`
	}
)
