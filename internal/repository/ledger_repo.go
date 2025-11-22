package repository

import (
	"context"
	"time"

	"github.com/ThePromisedNeverland/021trade/internal/logger"
	"github.com/ThePromisedNeverland/021trade/internal/models"
	"github.com/jmoiron/sqlx"
)

type ledgerRepo struct {
	db  *sqlx.DB
	log *logger.Logger
}

func NewLedgerRepo(db *sqlx.DB, log *logger.Logger) LedgerRepository {
	return &ledgerRepo{db: db, log: log}
}

func (l *ledgerRepo) AddEntry(ctx context.Context, entry models.LedgerEntry) error {
	if entry.CreatedAt.IsZero() {
		entry.CreatedAt = time.Now().UTC()
	}

	query := `
        INSERT INTO ledger
        (reward_id, symbol, quantity, inr_cost, brokerage_fee, stt_tax, gst_fee, other_fees, created_at)
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9);
    `

	_, err := l.db.ExecContext(ctx, query,
		entry.RewardID,
		entry.Symbol,
		entry.Quantity,
		entry.INRCost,
		entry.BrokerageFee,
		entry.STTTax,
		entry.GSTFee,
		entry.OtherFees,
		entry.CreatedAt,
	)

	if err != nil {
		l.log.WithError(err).Error("AddEntry failed")
	}

	return err
}

func (l *ledgerRepo) GetUserEntries(ctx context.Context, userID int64) ([]models.LedgerEntry, error) {
	query := `
		SELECT 
		    l.id,
		    l.reward_id,
		    l.symbol,
		    l.quantity,
		    l.inr_cost,
		    l.brokerage_fee,
		    l.stt_tax,
		    l.gst_fee,
		    l.other_fees,
		    l.created_at
		FROM ledger l
		JOIN rewards r ON l.reward_id = r.id
		WHERE r.user_id = $1
		ORDER BY l.created_at ASC;
	`

	var entries []models.LedgerEntry

	err := l.db.SelectContext(ctx, &entries, query, userID)
	if err != nil {
		l.log.WithError(err).Error("GetUserEntries failed")
		return nil, err
	}

	return entries, nil
}
