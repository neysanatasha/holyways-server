package Fundsdto

type FundCreateRequest struct {
	Title       string `json:"title" validate:"required"`
	Thumbnail   string `json:"thumbnail" form:"thumbnail" validate:"required"`
	Goal        string `json:"goals" validate:"required"`
	Description string `json:"description" validate:"required"`
	// UserID      int    `json:"user_id" form:"user_id"`k
}

type UpdateFundRequest struct {
	Title       string `json:"title" validate:"required"`
	Thumbnail   string `json:"thumbnail" validate:"required"`
	Goal        string `json:"goals" validate:"required"`
	Description string `json:"description" validate:"required"`
}
