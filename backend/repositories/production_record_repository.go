package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/lenardjombo/gallus/models"
)

type ProductionRecordRepository interface {
	Create(ctx context.Context, pr *models.ProductionRecord) error
	GetByID(ctx context.Context, id string) (*models.ProductionRecord, error)
	List(ctx context.Context) ([]*models.ProductionRecord, error)
	Update(ctx context.Context, pr *models.ProductionRecord) error
	Delete(ctx context.Context, id string) error
}

type PostgresProductionRecordRepository struct {
	Db *sql.DB
}

// Create implements ProductionRecordRepository
func (r *PostgresProductionRecordRepository) Create(ctx context.Context, pr *models.ProductionRecord) error {
	query := `
		INSERT INTO production_records (id, cohort_id, date, egg_count, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.Db.ExecContext(ctx, query,
		pr.ID, pr.CohortID, pr.Date, pr.EggCount, time.Now(), time.Now(),
	)
	return err
}

// GetByID implements ProductionRecordRepository
func (r *PostgresProductionRecordRepository) GetByID(ctx context.Context, id string) (*models.ProductionRecord, error) {
	query := `
		SELECT id, cohort_id, date, egg_count, created_at, updated_at
		FROM production_records
		WHERE id = $1
	`
	row := r.Db.QueryRowContext(ctx, query, id)

	var pr models.ProductionRecord
	err := row.Scan(&pr.ID, &pr.CohortID, &pr.Date, &pr.EggCount, &pr.CreatedAt, &pr.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &pr, err
}

// List implements ProductionRecordRepository
func (r *PostgresProductionRecordRepository) List(ctx context.Context) ([]*models.ProductionRecord, error) {
	query := `
		SELECT id, cohort_id, date, egg_count, created_at, updated_at
		FROM production_records
	`
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []*models.ProductionRecord
	for rows.Next() {
		var pr models.ProductionRecord
		err := rows.Scan(&pr.ID, &pr.CohortID, &pr.Date, &pr.EggCount, &pr.CreatedAt, &pr.UpdatedAt)
		if err != nil {
			return nil, err
		}
		records = append(records, &pr)
	}
	return records, nil
}

// Update implements ProductionRecordRepository
func (r *PostgresProductionRecordRepository) Update(ctx context.Context, pr *models.ProductionRecord) error {
	query := `
		UPDATE production_records
		SET cohort_id = $1, date = $2, egg_count = $3, updated_at = $4
		WHERE id = $5
	`
	_, err := r.Db.ExecContext(ctx, query,
		pr.CohortID, pr.Date, pr.EggCount, time.Now(), pr.ID,
	)
	return err
}

// Delete implements ProductionRecordRepository
func (r *PostgresProductionRecordRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM production_records WHERE id = $1`
	_, err := r.Db.ExecContext(ctx, query, id)
	return err
}

// Constructor
func NewProductionRecordRepository(db *sql.DB) ProductionRecordRepository {
	return &PostgresProductionRecordRepository{Db: db}
}
