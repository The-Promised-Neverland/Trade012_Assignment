package services

import (
	"context"

	"github.com/ThePromisedNeverland/021trade/internal/logger"
	"github.com/ThePromisedNeverland/021trade/internal/models"
	"github.com/ThePromisedNeverland/021trade/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
	log      *logger.Logger
}

func NewUserService(ur repository.UserRepository, log *logger.Logger) *UserService {
	return &UserService{
		userRepo: ur,
		log:      log,
	}
}

func (s *UserService) GetUser(ctx context.Context, userID int64) (*models.User, error) {
	return s.userRepo.GetUser(ctx, userID)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.userRepo.GetAllUsers(ctx)
}
