package services

import (
	"context"
	"math/rand"
	"time"

	"github.com/ThePromisedNeverland/021trade/internal/logger"
	"github.com/ThePromisedNeverland/021trade/internal/models"
	"github.com/ThePromisedNeverland/021trade/internal/repository"
)

type PriceService struct {
	priceRepo repository.PriceRepository
	log       *logger.Logger
}

func NewPriceService(pr repository.PriceRepository, log *logger.Logger) *PriceService {
	return &PriceService{
		priceRepo: pr,
		log:       log,
	}
}

func (s *PriceService) UpdatePrice(ctx context.Context, stock string) error {
	price := 100 + rand.Float64()*900

	err := s.priceRepo.InsertPrice(ctx, models.StockPrice{
		Symbol:         stock,
		Price:          price,
		PriceTimestamp: time.Now().UTC(),
	})
	if err != nil {
		s.log.WithError(err).Error("Failed to insert price")
	}
	return err
}

func (s *PriceService) GetLatestPrice(ctx context.Context, stock string) (float64, error) {
	price, err := s.priceRepo.GetLatestPrice(ctx, stock)
	if err != nil {
		s.log.WithError(err).Error("Failed to insert price")
	}
	return price, err
}
