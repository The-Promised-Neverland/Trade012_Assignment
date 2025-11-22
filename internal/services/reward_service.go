package services

import (
	"context"
	"time"

	"github.com/ThePromisedNeverland/021trade/internal/logger"
	"github.com/ThePromisedNeverland/021trade/internal/models"
	"github.com/ThePromisedNeverland/021trade/internal/repository"
)

type RewardService struct {
	rewardRepo repository.RewardRepository
	ledgerRepo repository.LedgerRepository
	priceRepo  repository.PriceRepository
	log        *logger.Logger
}

func NewRewardService(
	rr repository.RewardRepository,
	lr repository.LedgerRepository,
	pr repository.PriceRepository,
	log *logger.Logger,
) *RewardService {
	return &RewardService{
		rewardRepo: rr,
		ledgerRepo: lr,
		priceRepo:  pr,
		log:        log,
	}
}

func (s *RewardService) RewardUser(ctx context.Context, userID int64, symbol string, shares float64) (int64, error) {
	price, err := s.priceRepo.GetLatestPrice(ctx, symbol)
	if err != nil {
		s.log.Warn("Missing price — using 0 INR valuation")
		price = 0
	}

	timestamp := time.Now().UTC()

	rid, err := s.rewardRepo.Create(ctx, models.Reward{
		UserID:          userID,
		Symbol:          symbol,
		Quantity:        shares,
		BuyPrice:        price,
		RewardTimestamp: timestamp,
	})
	if err != nil {
		return 0, err
	}

	stockCost := price * shares
	brokerage := stockCost * 0.0003
	stt := stockCost * 0.001
	gst := brokerage * 0.18
	other := 0.0

	totalCost := stockCost + brokerage + stt + gst + other

	err = s.ledgerRepo.AddEntry(ctx, models.LedgerEntry{
		RewardID:     rid,
		Symbol:       symbol,
		Quantity:     shares,
		INRCost:      totalCost,
		BrokerageFee: brokerage,
		STTTax:       stt,
		GSTFee:       gst,
		OtherFees:    other,
	})
	if err != nil {
		s.log.Error("ledger entry insert failed")
	}

	return rid, nil
}

func (s *RewardService) GetTodayRewards(ctx context.Context, userID int64) ([]models.Reward, error) {
	return s.rewardRepo.GetTodayRewards(ctx, userID)
}

func (s *RewardService) GetHistoricalRecord(ctx context.Context, userID int64) (map[string]float64, error) {
	return s.rewardRepo.GetHistoricalReward(ctx, userID)
}

func (s *RewardService) GetStats(ctx context.Context, userID int64) (map[string]interface{}, error) {
	todayRewards, err := s.rewardRepo.GetTodayRewards(ctx, userID)
	if err != nil {
		return nil, err
	}

	todayMap := make(map[string]float64)
	for _, r := range todayRewards {
		todayMap[r.Symbol] += r.Quantity
	}

	allRewards, err := s.rewardRepo.GetAllRewards(ctx, userID)
	if err != nil {
		return nil, err
	}

	stockHoldings := make(map[string]float64)
	for _, r := range allRewards {
		stockHoldings[r.Symbol] += r.Quantity
	}

	portfolio := make(map[string]map[string]interface{})
	totalINR := 0.0

	for symbol, totalShares := range stockHoldings {
		if totalShares == 0 {
			continue
		}

		price, err := s.priceRepo.GetLatestPrice(ctx, symbol)
		if err != nil {
			s.log.Warnf("No price for %s, using 0", symbol)
			price = 0
		}

		value := price * totalShares
		portfolio[symbol] = map[string]interface{}{
			"shares":    totalShares,
			"price":     price,
			"value_inr": value,
		}
		totalINR += value
	}

	return map[string]interface{}{
		"today_shares":        todayMap,
		"total_portfolio_inr": totalINR,
	}, nil
}

func (s *RewardService) GetPortfolio(ctx context.Context, userID int64) (map[string]interface{}, error) {
	allRewards, err := s.rewardRepo.GetAllRewards(ctx, userID)
	if err != nil {
		return nil, err
	}

	stockHoldings := make(map[string]float64)
	for _, r := range allRewards {
		stockHoldings[r.Symbol] += r.Quantity
	}

	portfolio := make(map[string]map[string]interface{})
	totalINR := 0.0

	for symbol, totalShares := range stockHoldings {
		if totalShares == 0 {
			continue
		}

		price, err := s.priceRepo.GetLatestPrice(ctx, symbol)
		if err != nil {
			s.log.Warnf("No price for %s, using 0", symbol)
			price = 0
		}

		value := price * totalShares
		portfolio[symbol] = map[string]interface{}{
			"quantity":     totalShares,
			"latest_price": price,
			"value_inr":    value,
		}
		totalINR += value
	}

	return map[string]interface{}{
		"holdings":            portfolio,
		"total_portfolio_inr": totalINR,
	}, nil
}

func (s *RewardService) GetAllRewards(ctx context.Context, userID int64) ([]models.Reward, error) {
	return s.rewardRepo.GetAllRewards(ctx, userID)
}

func (s *RewardService) UpsertHistoryEntry(ctx context.Context, userID int64, date string, stock string, shares float64, value float64) error {
	return s.rewardRepo.UpsertHistoryEntry(ctx,userID,date,stock,shares,value)
}
