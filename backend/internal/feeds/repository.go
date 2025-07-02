package feeds

import (
	"context"
	"database/sql"
	"time"

	"github.com/lenardjombo/gallus/models"
)

type FeedsRepository interface {
	CreateFeedConsumption(ctx context.Context, c *models.FeedConsumption) error
	CreateFeedsPurchase(ctx context.Context, c *models.FeedPurchase) error
}

type PostgresFeedsRepository struct {
	DB *sql.DB
}

func (r *PostgresFeedsRepository) CreateFeedConsumption(ctx context.Context, c *models.FeedConsumption) error {
	query := `
		INSERT INTO feed_consumption (id, cohort_id, date, feed_kg, water_liters, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.DB.ExecContext(ctx, query,
		c.ID, c.CohortID, c.Date, c.FeedKG, c.WaterLiters, time.Now(),
	)
	return err
}

func (r *PostgresFeedsRepository) CreateFeedsPurchase(ctx context.Context, c *models.FeedPurchase) error {
	query := `
		INSERT INTO feed_purchases (id, supplier_id, purchase_date, cost, bags, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.DB.ExecContext(ctx, query,
		c.ID, c.SupplierID, time.Now(), c.Cost, c.Bags, time.Now(),
	)
	return err
}

func NewFeedsRepository(database *sql.DB) FeedsRepository {
	return &PostgresFeedsRepository{DB: database}
}
