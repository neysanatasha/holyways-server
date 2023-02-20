package Transactionsdto

import (
	"holyways/models"
	"time"
)

type TransactionResponse struct {
	ID           int                 `json:"id"`
	DonateAmount string              `json:"donateAmount" validate:"required"`
	Status       string              `json:"status" validate:"required"`
	CreatedAt    time.Time           `json:"startdate"`
	Fund         models.FundResponse `json:"fund" gorm:"foreignKey:fund_id"`
	// User         models.UserResponse `json:"user" gorm:"foreignKey:user_id"`
}
