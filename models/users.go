package models

// import "time"

type User struct {
	ID       int    `gorm:"primary_key:auto_increment" json:"id"`
	Fullname string `gorm:"type:varchar(300)" json:"fullName"`
	Email    string `gorm:"type:varchar(300)" json:"email"`
	Password string `gorm:"type:varchar(300)" json:"password"`
	Phone    string `gorm:"type:varchar(300)" json:"phone"`
	// TransctionID int                     `json:"transaction_id" form:"transaction_id"`
	Fund        []Fund        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"fund"`
	Transaction []Transaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"transaction"`
}

type UserResponse struct {
	ID       int    `json:"user_id"`
	Fullname string `json:"fullName"`
	Email    string `json:"email"`
}

func (UserResponse) TableName() string {
	return "users"
}
