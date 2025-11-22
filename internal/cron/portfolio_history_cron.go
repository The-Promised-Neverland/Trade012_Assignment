package cron

import (
	"context"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/ThePromisedNeverland/021trade/internal/logger"
	"github.com/ThePromisedNeverland/021trade/internal/services"
)

type PortfolioCron struct {
	cron          *cron.Cron
	userService   *services.UserService
	rewardService *services.RewardService
	priceService  *services.PriceService
	log           *logger.Logger
}

func NewPortfolioCron(u *services.UserService, r *services.RewardService, p *services.PriceService, log *logger.Logger) *PortfolioCron {
	return &PortfolioCron{
		cron:          cron.New(),
		userService:   u,
		rewardService: r,
		priceService:  p,
		log:           log,
	}
}

func (c *PortfolioCron) StartMidnightlyCron() {
	c.cron.AddFunc("0 0 * * *", func() {
		c.run()
	})
	c.cron.Start()
}

func (c *PortfolioCron) run() {
	ctx := context.Background()

	users, err := c.userService.GetAllUsers(ctx)
	if err != nil {
		c.log.WithError(err).Error("failed to fetch users")
		return
	}
	for _, u := range users {
		c.processUserPortfolio(ctx, u.ID)
	}
	c.log.Info("Cron Triggered. Portfolio updated")
}

func (c *PortfolioCron) processUserPortfolio(ctx context.Context, userID int64) {
	rewards, err := c.rewardService.GetAllRewards(ctx, userID)
	if err != nil {
		c.log.WithError(err).Error("failed to get rewards for user")
		return
	}
	shares := make(map[string]float64)
	for _, r := range rewards {
		shares[r.Symbol] += r.Quantity
	}
	today := time.Now().Format("2006-01-02")
	for symbol, totalShares := range shares {
		price, err := c.priceService.GetLatestPrice(ctx, symbol)
		if err != nil {
			price = 0
		}
		value := totalShares * price
		_ = c.rewardService.UpsertHistoryEntry(
			ctx, userID, today, symbol, totalShares, value,
		)
	}
}
