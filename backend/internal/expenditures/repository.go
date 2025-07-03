package expenditures

import (
	"context"
	"database/sql"
	"time"

	"github.com/lenardjombo/gallus/models"
)

type Expenditures interface {
	// Category CRUD
	CreateCategory(ctx context.Context, c *models.Category) error
	GetCategoryByID(ctx context.Context, id string) (*models.Category, error)
	ListCategories(ctx context.Context) ([]*models.Category, error)
	UpdateCategory(ctx context.Context, c *models.Category) error
	DeleteCategory(ctx context.Context, id string) error

	// Expenditure CRUD
	CreateExpenditure(ctx context.Context, e *models.Expenditure) error
	GetExpenditureByID(ctx context.Context, id string) (*models.Expenditure, error)
	ListExpenditures(ctx context.Context) ([]*models.Expenditure, error)
	UpdateExpenditure(ctx context.Context, e *models.Expenditure) error
	DeleteExpenditure(ctx context.Context, id string) error
}


type PostgresExpenditureRepository struct {
	DB *sql.DB
}

//Create a category
func (r *PostgresExpenditureRepository)CreateCategory(ctx context.Context, c *models.Category)error{
	query := `
		INSERT INTO categories (id, name, created_at)
		VALUES ($1, $2, $3)
	`
	_,err := r.DB.ExecContext(ctx,query,c.ID, c.Name, time.Now())
	return err
}

// Get category by ID
func (r *PostgresExpenditureRepository) GetCategoryByID(ctx context.Context, id string) (*models.Category, error) {
	query := `SELECT id, name, created_at FROM categories WHERE id = $1`
	row := r.DB.QueryRowContext(ctx, query, id)

	var c models.Category
	err := row.Scan(&c.ID, &c.Name, &c.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &c, nil
}


//Get list of categories
func (r *PostgresExpenditureRepository) ListCategories(ctx context.Context) ([]*models.Category, error) {
	query := `SELECT id, name, created_at FROM categories`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, &c)
	}
	return categories, nil
}

//Update a category
func (r *PostgresExpenditureRepository) UpdateCategory(ctx context.Context, c *models.Category) error {
	query := `UPDATE categories SET name = $1 WHERE id = $2`
	_, err := r.DB.ExecContext(ctx, query, c.Name, c.ID)
	return err
}

//Delete a category 
func (r *PostgresExpenditureRepository) DeleteCategory(ctx context.Context, id string) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}


//Create an expenditure 
func (r *PostgresExpenditureRepository)CreateExpenditure(ctx context.Context, c *models.Expenditure)error{
	query := `
		INSERT INTO expenditures (id, category_id, cohort_id, amount, name, purpose, date, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_,err := r.DB.ExecContext(ctx, query, c.ID, c.CategoryID, c.CohortID, c.Amount, c.Name, c.Purpose, time.Now(), time.Now())
	return err
}

// Get expenditure by ID
func (r *PostgresExpenditureRepository) GetExpenditureByID(ctx context.Context, id string) (*models.Expenditure, error) {
	query := `
		SELECT id, category_id, cohort_id, amount, name, purpose, date, created_at
		FROM expenditures WHERE id = $1
	`
	row := r.DB.QueryRowContext(ctx, query, id)

	var e models.Expenditure
	err := row.Scan(&e.ID, &e.CategoryID, &e.CohortID, &e.Amount, &e.Name, &e.Purpose, &e.Date, &e.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

//Get List of all expenditures
func (r *PostgresExpenditureRepository) ListExpenditures(ctx context.Context) ([]*models.Expenditure, error) {
	query := `
		SELECT id, category_id, cohort_id, amount, name, purpose, date, created_at
		FROM expenditures ORDER BY date DESC
	`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenditures []*models.Expenditure
	for rows.Next() {
		var e models.Expenditure
		if err := rows.Scan(&e.ID, &e.CategoryID, &e.CohortID, &e.Amount, &e.Name, &e.Purpose, &e.Date, &e.CreatedAt); err != nil {
			return nil, err
		}
		expenditures = append(expenditures, &e)
	}
	return expenditures, nil
}

//Update an expenditure 
func (r *PostgresExpenditureRepository) UpdateExpenditure(ctx context.Context, e *models.Expenditure) error {
	query := `
		UPDATE expenditures
		SET category_id = $1, cohort_id = $2, amount = $3, name = $4, purpose = $5, date = $6
		WHERE id = $7
	`
	_, err := r.DB.ExecContext(ctx, query, e.CategoryID, e.CohortID, e.Amount, e.Name, e.Purpose, e.Date, e.ID)
	return err
}

//Delete an expenditure 
func (r *PostgresExpenditureRepository) DeleteExpenditure(ctx context.Context, id string) error {
	query := `DELETE FROM expenditures WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}

// Constructor
func NewExpenditureRepository(database *sql.DB) Expenditures{
	return  &PostgresExpenditureRepository{DB: database}
}