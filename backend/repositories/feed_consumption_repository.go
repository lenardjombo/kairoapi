package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/lenardjombo/gallus/models"
)

type FeedConsumptionRepository interface {
	Create(ctx context.Context, fc *models.FeedConsumption) error
	GetByID(ctx context.Context, id string) (*models.FeedConsumption, error)
	List(ctx context.Context) ([]*models.FeedConsumption, error)
	Update(ctx context.Context, fc *models.FeedConsumption) error
	Delete(ctx context.Context, id string) error
}

type PostgresFeedConsumptionRepository struct {
	Db *sql.DB
}

// Create implements FeedConsumptionRepository
func (r *PostgresFeedConsumptionRepository) Create(ctx context.Context, fc *models.FeedConsumption) error {
	query := `
		INSERT INTO feed_consumption (id, cohort_id, date, feed_kg, water_liters, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.Db.ExecContext(ctx, query,
		fc.ID, fc.CohortID, fc.Date, fc.FeedKG, fc.WaterLiters, time.Now(),
	)
	return err
}

// GetByID implements FeedConsumptionRepository
func (r *PostgresFeedConsumptionRepository) GetByID(ctx context.Context, id string) (*models.FeedConsumption, error) {
	query := `
		SELECT id, cohort_id, date, feed_kg, water_liters, created_at
		FROM feed_consumption
		WHERE id = $1
	`
	row := r.Db.QueryRowContext(ctx, query, id)

	var fc models.FeedConsumption
	err := row.Scan(
		&fc.ID, &fc.CohortID, &fc.Date, &fc.FeedKG, &fc.WaterLiters, &fc.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &fc, err
}

// List implements FeedConsumptionRepository
func (r *PostgresFeedConsumptionRepository) List(ctx context.Context) ([]*models.FeedConsumption, error) {
	query := `
		SELECT id, cohort_id, date, feed_kg, water_liters, created_at
		FROM feed_consumption
	`
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*models.FeedConsumption
	for rows.Next() {
		var fc models.FeedConsumption
		err := rows.Scan(
			&fc.ID, &fc.CohortID, &fc.Date, &fc.FeedKG, &fc.WaterLiters, &fc.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, &fc)
	}
	return results, nil
}

// Update implements FeedConsumptionRepository
func (r *PostgresFeedConsumptionRepository) Update(ctx context.Context, fc *models.FeedConsumption) error {
	query := `
		UPDATE feed_consumption
		SET cohort_id = $1, date = $2, feed_kg = $3, water_liters = $4
		WHERE id = $5
	`
	_, err := r.Db.ExecContext(ctx, query,
		fc.CohortID, fc.Date, fc.FeedKG, fc.WaterLiters, fc.ID,
	)
	return err
}

// Delete implements FeedConsumptionRepository
func (r *PostgresFeedConsumptionRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM feed_consumption WHERE id = $1`
	_, err := r.Db.ExecContext(ctx, query, id)
	return err
}

// Constructor
func NewFeedConsumptionRepository(db *sql.DB) FeedConsumptionRepository {
	return &PostgresFeedConsumptionRepository{Db: db}
}
