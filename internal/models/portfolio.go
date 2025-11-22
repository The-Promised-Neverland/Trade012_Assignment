package models

import "time"

type PortfolioItem struct {
	ID        int64     `db:"id" json:"id"`
	Symbol    string    `db:"symbol" json:"symbol"`
	Price     float64   `db:"price" json:"price"`
	Timestamp time.Time `db:"timestamp" json:"timestamp"`
}

type Portfolio struct {
	UserID int64           `db:"user_id" json:"user_id"`
	Items  []PortfolioItem `json:"items"`
	Total  float64         `json:"total"`
}
