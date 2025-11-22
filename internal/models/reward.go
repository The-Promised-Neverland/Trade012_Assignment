package models

import "time"

type Reward struct {
	ID              int64     `db:"id" json:"id"`
	UserID          int64     `db:"user_id" json:"user_id"`
	Symbol          string    `db:"symbol" json:"symbol"`
	Quantity        float64   `db:"quantity" json:"quantity"`
	BuyPrice        float64   `db:"buy_price" json:"buy_price"`
	RewardTimestamp time.Time `db:"reward_timestamp" json:"reward_timestamp"`
}
