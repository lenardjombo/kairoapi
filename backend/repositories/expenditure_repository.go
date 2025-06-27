package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/lenardjombo/gallus/models"
)

type ExpenditureRepository interface {
	Create(ctx context.Context, e *models.Expenditure) error
	GetByID(ctx context.Context, id string) (*models.Expenditure, error)
	List(ctx context.Context) ([]*models.Expenditure, error)
	Update(ctx context.Context, e *models.Expenditure) error
	Delete(ctx context.Context, id string) error
}

type PostgresExpenditureRepository struct {
	Db *sql.DB
}

// Create implements ExpenditureRepository
func (r *PostgresExpenditureRepository) Create(ctx context.Context, e *models.Expenditure) error {
	query := `
		INSERT INTO expenditures (id, category_id, cohort_id, amount, name, purpose, date, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.Db.ExecContext(ctx, query,
		e.ID, e.CategoryID, e.CohortID, e.Amount, e.Name, e.Purpose, e.Date, time.Now(),
	)
	return err
}

// GetByID implements ExpenditureRepository
func (r *PostgresExpenditureRepository) GetByID(ctx context.Context, id string) (*models.Expenditure, error) {
	query := `
		SELECT id, category_id, cohort_id, amount, name, purpose, date, created_at
		FROM expenditures
		WHERE id = $1
	`
	row := r.Db.QueryRowContext(ctx, query, id)

	var exp models.Expenditure
	err := row.Scan(
		&exp.ID, &exp.CategoryID, &exp.CohortID, &exp.Amount,
		&exp.Name, &exp.Purpose, &exp.Date, &exp.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &exp, err
}

// List implements ExpenditureRepository
func (r *PostgresExpenditureRepository) List(ctx context.Context) ([]*models.Expenditure, error) {
	query := `
		SELECT id, category_id, cohort_id, amount, name, purpose, date, created_at
		FROM expenditures
	`
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenditures []*models.Expenditure
	for rows.Next() {
		var exp models.Expenditure
		err := rows.Scan(
			&exp.ID, &exp.CategoryID, &exp.CohortID, &exp.Amount,
			&exp.Name, &exp.Purpose, &exp.Date, &exp.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		expenditures = append(expenditures, &exp)
	}
	return expenditures, nil
}

// Update implements ExpenditureRepository
func (r *PostgresExpenditureRepository) Update(ctx context.Context, e *models.Expenditure) error {
	query := `
		UPDATE expenditures
		SET category_id = $1, cohort_id = $2, amount = $3, name = $4,
		    purpose = $5, date = $6
		WHERE id = $7
	`
	_, err := r.Db.ExecContext(ctx, query,
		e.CategoryID, e.CohortID, e.Amount, e.Name,
		e.Purpose, e.Date, e.ID,
	)
	return err
}

// Delete implements ExpenditureRepository
func (r *PostgresExpenditureRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM expenditures WHERE id = $1`
	_, err := r.Db.ExecContext(ctx, query, id)
	return err
}

// Constructor
func NewExpenditureRepository(db *sql.DB) ExpenditureRepository {
	return &PostgresExpenditureRepository{Db: db}
}
