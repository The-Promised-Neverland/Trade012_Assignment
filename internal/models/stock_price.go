package models

import (
	"time"
)

type StockPrice struct {
	Symbol         string    `db:"symbol" json:"symbol"`
	Price          float64   `db:"price" json:"price"`
	PriceTimestamp time.Time `db:"price_timestamp" json:"price_timestamp"`
}
