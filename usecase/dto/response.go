package dto

import "time"

type (
	ChargeWalletResponse struct {
		ID      string `json:"id"`
		Balance int    `json:"balance"`
	}

	Discount struct {
		Amount    int
		CreatedAt time.Time
	}

	GetDiscountsByIDResponse struct {
		Discounts []Discount `json:"discounts"`
	}
)
