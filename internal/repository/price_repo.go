package repository

import (
	"context"
	"errors"
	"time"

	"github.com/ThePromisedNeverland/021trade/internal/logger"
	"github.com/ThePromisedNeverland/021trade/internal/models"
	"github.com/jmoiron/sqlx"
)

type priceRepo struct {
	db  *sqlx.DB
	log *logger.Logger
}

func NewPriceRepo(db *sqlx.DB, log *logger.Logger) PriceRepository {
	return &priceRepo{
		db:  db,
		log: log,
	}
}

func (p *priceRepo) InsertPrice(ctx context.Context, stockPrice models.StockPrice) error {
	if stockPrice.PriceTimestamp.IsZero() {
		stockPrice.PriceTimestamp = time.Now().UTC()
	}
	query := `
		INSERT INTO stock_prices (symbol, price, price_timestamp)
		VALUES ($1, $2, NOW())
		ON CONFLICT (symbol)
		DO UPDATE SET
			price = EXCLUDED.price,
			price_timestamp = NOW();
	`

	_, err := p.db.ExecContext(ctx, query,
		stockPrice.Symbol,
		stockPrice.Price,
	)

	if err != nil {
		p.log.WithError(err).Error("Insert Price failed")
		return err
	}

	return nil
}

func (p *priceRepo) GetLatestPrice(ctx context.Context, symbol string) (float64, error) {
	query := `
		SELECT price
		FROM stock_prices 
		WHERE symbol = $1
		ORDER BY price_timestamp DESC
		LIMIT 1;
	`
	var price float64
	err := p.db.GetContext(ctx, &price, query, symbol)
	if err != nil {
		p.log.WithError(err).Warnf("no latest price for stock=%s", symbol)
		return 0, errors.New("no price found")
	}

	return price, nil
}
