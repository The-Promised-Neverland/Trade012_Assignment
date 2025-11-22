package models

import "time"

type Stock struct {
	Symbol      string    `db:"symbol" json:"symbol"`
	CompanyName string    `db:"company_name" json:"company_name"`
	ISIN        string    `db:"isin" json:"isin"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
