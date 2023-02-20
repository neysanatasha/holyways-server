package Usersdto

type CreateUserRequest struct {
	Fullname string `json:"fullName" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}

type UpdateUserRequest struct {
	Fullname string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
