package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/lenardjombo/gallus/models"
)

type SupplierRepository interface {
	Create(ctx context.Context, s *models.Supplier) error
	GetByID(ctx context.Context, id string) (*models.Supplier, error)
	List(ctx context.Context) ([]*models.Supplier, error)
	Update(ctx context.Context, s *models.Supplier) error
	Delete(ctx context.Context, id string) error
}

type PostgresSupplierRepository struct {
	Db *sql.DB
}

// Create implements SupplierRepository
func (r *PostgresSupplierRepository) Create(ctx context.Context, s *models.Supplier) error {
	query := `
		INSERT INTO suppliers (id, name, contact, created_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.Db.ExecContext(ctx, query,
		s.ID, s.Name, s.Contact, time.Now(),
	)
	return err
}

// GetByID implements SupplierRepository
func (r *PostgresSupplierRepository) GetByID(ctx context.Context, id string) (*models.Supplier, error) {
	query := `
		SELECT id, name, contact, created_at
		FROM suppliers
		WHERE id = $1
	`
	row := r.Db.QueryRowContext(ctx, query, id)

	var s models.Supplier
	err := row.Scan(&s.ID, &s.Name, &s.Contact, &s.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &s, err
}

// List implements SupplierRepository
func (r *PostgresSupplierRepository) List(ctx context.Context) ([]*models.Supplier, error) {
	query := `
		SELECT id, name, contact, created_at
		FROM suppliers
	`
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suppliers []*models.Supplier
	for rows.Next() {
		var s models.Supplier
		err := rows.Scan(&s.ID, &s.Name, &s.Contact, &s.CreatedAt)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, &s)
	}
	return suppliers, nil
}

// Update implements SupplierRepository
func (r *PostgresSupplierRepository) Update(ctx context.Context, s *models.Supplier) error {
	query := `
		UPDATE suppliers
		SET name = $1, contact = $2
		WHERE id = $3
	`
	_, err := r.Db.ExecContext(ctx, query, s.Name, s.Contact, s.ID)
	return err
}

// Delete implements SupplierRepository
func (r *PostgresSupplierRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM suppliers WHERE id = $1`
	_, err := r.Db.ExecContext(ctx, query, id)
	return err
}

// Constructor
func NewSupplierRepository(db *sql.DB) SupplierRepository {
	return &PostgresSupplierRepository{Db: db}
}
