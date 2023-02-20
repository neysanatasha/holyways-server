package Transactionsdto

import "time"

type TransactionCreateRequest struct {
	FundID       int       `json:"fund_id" form:"fund_id"`
	Status       string    `json:"status"`
	UserID       int       `json:"user_id" form:"user_id"`
	CreatedAt    time.Time `json:"startdate"`
	DonateAmount string    `json:"donateAmount" form:"donateAmount"`
}
type UpdateTransactionRequest struct {
	DonateAmount string `json:"donateAmount" validate:"required"`
	Status       string `json:"status" validate:"required"`
	UserID       int    `json:"user_id" form:"user_id"`
}
