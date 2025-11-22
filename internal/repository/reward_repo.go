package repository

import (
	"context"

	"github.com/ThePromisedNeverland/021trade/internal/logger"
	"github.com/ThePromisedNeverland/021trade/internal/models"
	"github.com/jmoiron/sqlx"
)

type rewardRepo struct {
	db  *sqlx.DB
	log *logger.Logger
}

func NewRewardRepo(db *sqlx.DB, log *logger.Logger) RewardRepository {
	return &rewardRepo{db: db, log: log}
}

func (r *rewardRepo) Create(ctx context.Context, reward models.Reward) (int64, error) {
	query := `
		INSERT INTO rewards (user_id, symbol, quantity, buy_price, reward_timestamp)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	var id int64

	err := r.db.QueryRowxContext(ctx, query,
		reward.UserID,
		reward.Symbol,
		reward.Quantity,
		reward.BuyPrice,
		reward.RewardTimestamp,
	).Scan(&id)

	if err != nil {
		r.log.WithError(err).Error("reward create failed")
		return 0, err
	}

	return id, nil
}

func (r *rewardRepo) GetTodayRewards(ctx context.Context, userID int64) ([]models.Reward, error) {
	query := `
		SELECT id, user_id, symbol, quantity, buy_price, reward_timestamp
		FROM rewards
		WHERE user_id = $1
		  AND DATE(reward_timestamp) = CURRENT_DATE
		ORDER BY reward_timestamp ASC;
	`

	out := []models.Reward{}
	err := r.db.SelectContext(ctx, &out, query, userID)
	if err != nil {
		r.log.WithError(err).Error("GetTodayRewards failed")
		return nil, err
	}

	return out, nil
}

func (r *rewardRepo) GetAllRewards(ctx context.Context, userID int64) ([]models.Reward, error) {
	query := `
		SELECT id, user_id, symbol, quantity, reward_timestamp
		FROM rewards
		WHERE user_id = $1
		ORDER BY reward_timestamp ASC;
	`

	var out []models.Reward
	err := r.db.SelectContext(ctx, &out, query, userID)
	if err != nil {
		r.log.WithError(err).Error("GetAllRewards failed")
		return nil, err
	}

	return out, nil
}

func (r *rewardRepo) GetHistoricalReward(ctx context.Context, userID int64) (map[string]float64, error) {
	query := `
		SELECT 
			TO_CHAR(date, 'YYYY-MM-DD') AS date, 
			SUM(value_inr) AS total_inr
		FROM user_portfolio_history
		WHERE user_id = $1
		GROUP BY date
		ORDER BY date ASC;
	`

	type row struct {
		Date     string  `db:"date"`
		TotalINR float64 `db:"total_inr"`
	}

	var records []row

	err := r.db.SelectContext(ctx, &records, query, userID)
	if err != nil {
		r.log.WithError(err).Error("GetHistoricalRecordByUserID failed")
		return nil, err
	}

	out := make(map[string]float64)

	for _, rec := range records {
		out[rec.Date] = rec.TotalINR
	}

	return out, nil
}

func (r *rewardRepo) UpsertHistoryEntry(ctx context.Context, userID int64, date string, stock string, shares float64, value float64) error {
	query := `
        INSERT INTO user_portfolio_history (user_id, date, symbol, total_shares, value_inr)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (user_id, date, symbol)
        DO UPDATE SET
            total_shares = EXCLUDED.total_shares,
            value_inr = EXCLUDED.value_inr;
    `

	_, err := r.db.ExecContext(ctx, query,
		userID,
		date,
		stock,
		shares,
		value,
	)

	if err != nil {
		r.log.WithError(err).Error("UpsertHistoryEntry failed")
		return err
	}

	return nil
}
