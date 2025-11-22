package models

import "time"

type UserPortfolioHistory struct {
	UserID      int64     `db:"user_id" json:"user_id"`
	Date        time.Time `db:"date" json:"date"`
	Symbol      string    `db:"symbol" json:"symbol"`
	TotalShares float64   `db:"total_shares" json:"total_shares"`
	ValueINR    float64   `db:"value_inr" json:"value_inr"`
}
