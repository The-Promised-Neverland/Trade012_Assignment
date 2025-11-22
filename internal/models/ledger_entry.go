package models

import "time"

type LedgerEntry struct {
    ID           int64     `db:"id" json:"id"`
    RewardID     int64     `db:"reward_id" json:"reward_id"`
    Symbol       string    `db:"symbol" json:"symbol"`
    Quantity     float64   `db:"quantity" json:"quantity"`
    INRCost      float64   `db:"inr_cost" json:"inr_cost"`
    BrokerageFee float64   `db:"brokerage_fee" json:"brokerage_fee"`
    STTTax       float64   `db:"stt_tax" json:"stt_tax"`
    GSTFee       float64   `db:"gst_fee" json:"gst_fee"`
    OtherFees    float64   `db:"other_fees" json:"other_fees"`
    CreatedAt    time.Time `db:"created_at" json:"created_at"`
}
