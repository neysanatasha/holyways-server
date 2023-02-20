package models

import "time"

type Transaction struct {
	ID           int          `gorm:"primary_key:auto_increment" json:"id"`
	DonateAmount string       `gorm:"type:varchar(300)" json:"donateAmount" form:"donateAmount"`
	Status       string       `gorm:"type:varchar(300)" json:"status" form:"status"`
	CreatedAt    time.Time    `json:"startdate"`
	UserID       int          `json:"user_id" form:"user_id"`
	User         UserResponse `gorm:"foreignKey:UserID" json:"user"`
	FundID       int          `json:"fund_id"`
	Fund         FundResponse `json:"fund" `
}
