package models

type CreateRewardRequest struct {
	UserID int64   `json:"user_id" binding:"required"`
	Stock  string  `json:"stock" binding:"required"`
	Shares float64 `json:"shares" binding:"required"`
}
