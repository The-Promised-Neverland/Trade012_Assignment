package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/ThePromisedNeverland/021trade/internal/logger"
	"github.com/ThePromisedNeverland/021trade/internal/models"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db  *sqlx.DB
	log *logger.Logger
}

func NewUserRepo(db *sqlx.DB, log *logger.Logger) UserRepository {
	return &userRepo{db: db, log: log}
}

func (u *userRepo) GetUser(ctx context.Context, id int64) (*models.User, error) {
	query := `
        SELECT id, name, email, created_at
        FROM users
        WHERE id = $1
        LIMIT 1;
    `

	var usr models.User
	err := u.db.GetContext(ctx, &usr, query, id)
	fmt.Println(usr, err)
	if err != nil {
		u.log.WithError(err).Warnf("user not found id=%d", id)
		return nil, errors.New("user not found")
	}

	return &usr, nil
}

func (u *userRepo) GetAllUsers(ctx context.Context) ([]models.User, error) {
	query := `
		SELECT id, name, email, created_at
		FROM users
		ORDER BY id ASC;
	`

	var users []models.User

	err := u.db.SelectContext(ctx, &users, query)
	if err != nil {
		u.log.WithError(err).Error("GetAllUsers failed")
		return nil, err
	}

	return users, nil
}
