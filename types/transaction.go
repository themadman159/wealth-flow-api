package types

import "time"

type TransactionRequest struct {
	Type            string    `json:"type" validate:"required"`
	Category        string    `json:"category"`
	Amount          int64     `json:"amount" validate:"required"`
	Description     string    `json:"description"`
	TransactionDate time.Time `json:"transaction_date" validate:"required"`
}

type TransactionGetAll struct {
	ID              uint      `json:"id"`
	Type            string    `json:"type"`
	Category        string    `json:"category"`
	Amount          int64     `json:"amount"`
	Description     string    `json:"description"`
	TransactionDate time.Time `json:"transaction_date"`
}

type TransactionGetAllResponse struct {
	Transaction []TransactionGetAll `json:"transaction"`
	Total       int64               `json:"total"`
}
