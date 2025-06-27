package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/lenardjombo/gallus/models"
)

type CategoryRepository interface {
	Create(ctx context.Context, c *models.Category) error
	GetByID(ctx context.Context, id string) (*models.Category, error)
	List(ctx context.Context) ([]*models.Category, error)
	Update(ctx context.Context, c *models.Category) error
	Delete(ctx context.Context, id string) error
}

type PostgresCategoryRepository struct {
	Db *sql.DB
}

// Create implements CategoryRepository
func (r *PostgresCategoryRepository) Create(ctx context.Context, c *models.Category) error {
	query := `
		INSERT INTO categories (id, name, created_at)
		VALUES ($1, $2, $3)
	`
	_, err := r.Db.ExecContext(ctx, query, c.ID, c.Name, time.Now())
	return err
}

// GetByID implements CategoryRepository
func (r *PostgresCategoryRepository) GetByID(ctx context.Context, id string) (*models.Category, error) {
	query := `SELECT id, name, created_at FROM categories WHERE id = $1`
	row := r.Db.QueryRowContext(ctx, query, id)

	var cat models.Category
	err := row.Scan(&cat.ID, &cat.Name, &cat.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &cat, err
}

// List implements CategoryRepository
func (r *PostgresCategoryRepository) List(ctx context.Context) ([]*models.Category, error) {
	query := `SELECT id, name, created_at FROM categories`
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.Category
	for rows.Next() {
		var cat models.Category
		err := rows.Scan(&cat.ID, &cat.Name, &cat.CreatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &cat)
	}
	return categories, nil
}

// Update implements CategoryRepository
func (r *PostgresCategoryRepository) Update(ctx context.Context, c *models.Category) error {
	query := `
		UPDATE categories
		SET name = $1
		WHERE id = $2
	`
	_, err := r.Db.ExecContext(ctx, query, c.Name, c.ID)
	return err
}

// Delete implements CategoryRepository
func (r *PostgresCategoryRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := r.Db.ExecContext(ctx, query, id)
	return err
}

// Constructor
func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &PostgresCategoryRepository{Db: db}
}
