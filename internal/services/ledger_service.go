package services

import (
	"context"

	"github.com/ThePromisedNeverland/021trade/internal/logger"
	"github.com/ThePromisedNeverland/021trade/internal/models"
	"github.com/ThePromisedNeverland/021trade/internal/repository"
)

type LedgerService struct {
	ledgerRepo repository.LedgerRepository
	log        *logger.Logger
}

func NewLedgerService(lr repository.LedgerRepository, log *logger.Logger) *LedgerService {
	return &LedgerService{
		ledgerRepo: lr,
		log:        log,
	}
}

func (s *LedgerService) AddEntry(ctx context.Context, entry models.LedgerEntry) error {
	return s.ledgerRepo.AddEntry(ctx, entry)
}

func (s *LedgerService) GetUserEntries(ctx context.Context, userID int64) ([]models.LedgerEntry, error) {
	return s.ledgerRepo.GetUserEntries(ctx, userID)
}
