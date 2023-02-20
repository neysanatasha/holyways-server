package authdto

type RegisterResponse struct {
	Fullname string `gorm:"type:varchar(300)" json:"fullName"`
	Email    string `gorm:"type:varchar(300)" json:"email"`
	Password string `gorm:"type:varchar(300)" json:"password"`
}

type LoginResponse struct {
	Fullname string `gorm:"type:varchar(300)" json:"fullName"`
	Email    string `gorm:"type:varchar(300)" json:"email"`
	Token    string `gorm:"type:varchar(300)" json:"token"`
	Phone    string `gorm:"type:varchar(300)" json:"phone"`
}

type CheckAuthResponse struct {
	ID       int    `gorm:"type:int" json:"id"`
	Fullname string `gorm:"type:varchar(300)" json:"fullName"`
	Email    string `gorm:"type:varchar(300)" json:"email"`
	Token    string `gorm:"type:varchar(300)" json:"token"`
	Phone    string `gorm:"type:varchar(300)" json:"phone"`
}
