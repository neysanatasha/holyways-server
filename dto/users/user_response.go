package Usersdto

import "holyways/models"

type UserResponse struct {
	ID          int                `json:"id"`
	Fullname    string             `json:"fullName" validate:"required"`
	Email       string             `json:"email" validate:"required"`
	Phone       string             `json:"phone" validate:"required"`
	Fund        models.Fund        `json:"fund" gorm:"foreignKey:fund_id"`
	Transaction models.Transaction `json:"transaction" gorm:"foreignKey:transaction_id"`
}
