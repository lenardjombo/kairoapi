package expenditures

import (
	"context"
	"database/sql"
	"time"

	"github.com/lenardjombo/gallus/models"
)

type Expenditures interface {
	CreateCategory(ctx context.Context, c *models.Category)error
	CreateExpenditure(ctx context.Context, c *models.Expenditure)error
}

type PostgresExpenditureRepository struct {
	DB *sql.DB
}

func (r *PostgresExpenditureRepository)CreateCategory(ctx context.Context, c *models.Category)error{
	query := `
		INSERT INTO categories (id, name, created_at)
		VALUES ($1, $2, $3)
	`
	_,err := r.DB.ExecContext(ctx,query,c.ID, c.Name, time.Now())
	return err
}

func (r *PostgresExpenditureRepository)CreateExpenditure(ctx context.Context, c *models.Expenditure)error{
	query := `
		INSERT INTO expenditures (id, category_id, cohort_id, amount, name, purpose, date, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_,err := r.DB.ExecContext(ctx, query, c.ID, c.CategoryID, c.CohortID, c.Amount, c.Name, c.Purpose, time.Now(), time.Now())
	return err
}

// Constructor
func NewExpenditureRepository(database *sql.DB) Expenditures{
	return  &PostgresExpenditureRepository{DB: database}
}