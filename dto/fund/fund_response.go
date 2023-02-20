package Fundsdto

type FundResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title" validate:"required"`
	Thumbnail   string `json:"thumbnail" validate:"required"`
	Goal        string `json:"goal" validate:"required"`
	Description string `json:"description" validate:"required"`
}
