package repository

import (
	"context"

	"github.com/ThePromisedNeverland/021trade/internal/models"
)

type RewardRepository interface {
	Create(ctx context.Context, reward models.Reward) (int64, error)
	GetTodayRewards(ctx context.Context, userID int64) ([]models.Reward, error)
	GetAllRewards(ctx context.Context, userID int64) ([]models.Reward, error)
	GetHistoricalReward(ctx context.Context, userID int64) (map[string]float64, error)
	UpsertHistoryEntry(ctx context.Context, userID int64, date string, stock string, shares float64, value float64) error
}

type LedgerRepository interface {
	AddEntry(ctx context.Context, entry models.LedgerEntry) error
	GetUserEntries(ctx context.Context, userID int64) ([]models.LedgerEntry, error)
}

type PriceRepository interface {
	InsertPrice(ctx context.Context, p models.StockPrice) error
	GetLatestPrice(ctx context.Context, stock string) (float64, error)
}

type UserRepository interface {
	GetUser(ctx context.Context, id int64) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
}
