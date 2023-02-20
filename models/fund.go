package models

type Fund struct {
	ID          int          `gorm:"primary_key:auto_increment" json:"id"`
	Title       string       `gorm:"type:varchar(300)" json:"title" form:"title"`
	Thumbnail   string       `gorm:"type:varchar(300)" json:"thumbnail" form:"image"`
	Goal        string       `gorm:"type:varchar(300)" json:"goals" form:"goals"`
	Description string       `gorm:"type:text" json:"description" form:"description"`
	UserID      int          `json:"user_id" form:"user_id"`
	User        UserResponse `json:"user"`
	// TransactionID int                     `json:"transaction_id" form:"transaction_id"`
	Transaction []Transaction `json:"transaction"`
}

type FundResponse struct {
	ID          int    `json:"fund_id" form:"fund_id" gorm:"primary_key:auto_increment"`
	Title       string `json:"title"`
	Thumbnail   string `json:"thumbnail"`
	Goal        string `json:"goal"`
	Description string `json:"description"`
	// UserID      int          `json:"user_id"`
	// User        UserResponse `json:"user"`
}

func (FundResponse) TableName() string {
	return "funds"
}
