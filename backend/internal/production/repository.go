// Data Access layer for Production
package production

import (
	"context"
	"database/sql"
	"time"

	"github.com/lenardjombo/gallus/models"
)

// Interface - Defines the methods for interacting with Production data
type ProductionRepository interface {
	Create(ctx context.Context, c *models.ProductionRecord) error
	GetByID(ctx context.Context, id string) (*models.ProductionRecord, error)
	List(ctx context.Context) ([]*models.ProductionRecord, error)
	Update(ctx context.Context, c *models.ProductionRecord) error
	Delete(ctx context.Context, id string) error
}

// ProductionRepo implements the ProductionRepository interface
type ProductionRepo struct {
	DB *sql.DB
}

// Create adds a new production record to the database
func (r *ProductionRepo) Create(ctx context.Context, c *models.ProductionRecord) error {
	query := `
		INSERT INTO production_records (id, cohort_id, date, egg_count, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.DB.ExecContext(ctx, query,
		c.ID,
		c.CohortID,
		c.Date,
		c.EggCount,
		time.Now(), // created_at
		time.Now(), // updated_at
	)
	return err
}

// GetByID retrieves a single production record by ID
func (r *ProductionRepo) GetByID(ctx context.Context, id string) (*models.ProductionRecord, error) {
	query := `
		SELECT id, cohort_id, date, egg_count, created_at, updated_at
		FROM production_records
		WHERE id = $1
	`
	row := r.DB.QueryRowContext(ctx, query, id)

	var record models.ProductionRecord
	err := row.Scan(
		&record.ID,
		&record.CohortID,
		&record.Date,
		&record.EggCount,
		&record.CreatedAt,
		&record.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// List returns all production records ordered by date descending
func (r *ProductionRepo) List(ctx context.Context) ([]*models.ProductionRecord, error) {
	query := `
		SELECT id, cohort_id, date, egg_count, created_at, updated_at
		FROM production_records
		ORDER BY date DESC
	`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []*models.ProductionRecord
	for rows.Next() {
		var record models.ProductionRecord
		err := rows.Scan(
			&record.ID,
			&record.CohortID,
			&record.Date,
			&record.EggCount,
			&record.CreatedAt,
			&record.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		records = append(records, &record)
	}
	return records, nil
}

// Update modifies an existing production record
func (r *ProductionRepo) Update(ctx context.Context, c *models.ProductionRecord) error {
	query := `
		UPDATE production_records
		SET cohort_id = $1, date = $2, egg_count = $3, updated_at = $4
		WHERE id = $5
	`
	_, err := r.DB.ExecContext(ctx, query,
		c.CohortID,
		c.Date,
		c.EggCount,
		time.Now(),
		c.ID,
	)
	return err
}

// Delete removes a production record by ID
func (r *ProductionRepo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM production_records WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}

// Constructor - Returns a new instance of ProductionRepo with the provided database
func NewProductionRepository(database *sql.DB) ProductionRepository {
	return &ProductionRepo{DB: database}
}
