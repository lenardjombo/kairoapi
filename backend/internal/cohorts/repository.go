package cohorts

import (
	"context"
	"database/sql"
	"time"

	"github.com/lenardjombo/gallus/models"
)

type CohortRepository interface {
	Create(ctx context.Context, c *models.Cohort) error
	GetByID(ctx context.Context, id string) (*models.Cohort, error)
	List(ctx context.Context) ([]*models.Cohort, error)
	Update(ctx context.Context, c *models.Cohort) error
	Delete(ctx context.Context, id string) error
}

type PostgresCohortRepository struct {
	DB *sql.DB
}

// Create implements CohortRepository.
func (r *PostgresCohortRepository) Create(ctx context.Context, c *models.Cohort) error {
    query := `
        INSERT INTO cohorts (id, name, breed, start_date, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
    _, err := r.DB.ExecContext(ctx, query,
        c.ID, c.Name, c.StartDate, time.Now(), time.Now(),
    )
    return err
}

// GetById implements CohortRepository.
func (r *PostgresCohortRepository) GetByID(ctx context.Context, id string) (*models.Cohort, error) {
    query := `SELECT id, name, breed, start_date, created_at, updated_at FROM cohorts WHERE id = $1`
    row := r.DB.QueryRowContext(ctx, query, id)

    var c models.Cohort
    err := row.Scan(&c.ID, &c.Name,  &c.StartDate, &c.CreatedAt, &c.UpdatedAt)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    return &c, err
}

// List implements CohortRepository.
func (r *PostgresCohortRepository) List(ctx context.Context) ([]*models.Cohort, error) {
    query := `SELECT id, name, breed, start_date, created_at, updated_at FROM cohorts`
    rows, err := r.DB.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var cohorts []*models.Cohort
    for rows.Next() {
        var c models.Cohort
        err := rows.Scan(&c.ID, &c.Name, &c.StartDate, &c.CreatedAt, &c.UpdatedAt)
        if err != nil {
            return nil, err
        }
        cohorts = append(cohorts, &c)
    }
    return cohorts, nil
}


// Update implements CohortRepository.
func (r *PostgresCohortRepository) Update(ctx context.Context, c *models.Cohort) error {
    query := `
        UPDATE cohorts
        SET name = $1, breed = $2, start_date = $3, updated_at = $4
        WHERE id = $5
    `
    _, err := r.DB.ExecContext(ctx, query, c.Name, c.StartDate, time.Now(), c.ID)
    return err
}

// Delete implements CohortRepository.
func (r *PostgresCohortRepository) Delete(ctx context.Context, id string) error {
    query := `DELETE FROM cohorts WHERE id = $1`
    _, err := r.DB.ExecContext(ctx, query, id)
    return err
}


func NewCohortRepository(db *sql.DB) CohortRepository {
	return &PostgresCohortRepository{DB: db}
}
