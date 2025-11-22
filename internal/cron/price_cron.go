package cron

import (
	"context"
	"math/rand"
	"time"

	"github.com/ThePromisedNeverland/021trade/internal/logger"
	"github.com/ThePromisedNeverland/021trade/internal/models"
	"github.com/ThePromisedNeverland/021trade/internal/repository"
	"github.com/robfig/cron/v3"
)

var TrackedStocks = []string{
	"RELIANCE",
	"TCS",
	"INFOSYS",
	"HDFC",
	"ICICIBANK",
}

type PriceUpdater struct {
	priceRepo repository.PriceRepository
	log       *logger.Logger
	cron      *cron.Cron
}

func NewPriceUpdater(pr repository.PriceRepository, log *logger.Logger) *PriceUpdater {
	return &PriceUpdater{
		priceRepo: pr,
		log:       log,
		cron:      cron.New(cron.WithSeconds()),
	}
}

func (p *PriceUpdater) GenerateRandomPrice(stock string) float64 {
	basePrice := map[string]float64{
		"RELIANCE":  2500,
		"TCS":       3500,
		"INFOSYS":   1500,
		"HDFC":      1600,
		"ICICIBANK": 1000,
	}
	base, exists := basePrice[stock]
	if !exists {
		base = 1000
	}
	variance := base * 0.1
	price := base + (rand.Float64()*2-1)*variance
	return price
}

func (p *PriceUpdater) UpdateAllPrices() {
	for _, stock := range TrackedStocks {
		price := p.GenerateRandomPrice(stock)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := p.priceRepo.InsertPrice(ctx, models.StockPrice{
			Symbol:         stock,
			Price:          price,
			PriceTimestamp: time.Now().UTC(),
		})
		if err != nil {
			p.log.WithField("stock", stock).Errorf("Failed to update price: %v", err)
			continue
		}
	}
	p.log.Info("Cron Triggered. Price updated")
}

func (p *PriceUpdater) StartHourlyCron() {
	p.UpdateAllPrices()
	_, err := p.cron.AddFunc("*/10 * * * * *", func() {
		p.UpdateAllPrices()
	})
	if err != nil {
		p.log.Errorf("Failed to schedule cron job: %v", err)
		return
	}
	p.cron.Start()
}

func (p *PriceUpdater) Stop() {
	p.cron.Stop()
	p.log.Info("Price updater cron stopped")
}
