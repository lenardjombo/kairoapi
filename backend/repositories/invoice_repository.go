package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/lenardjombo/gallus/models"
)

type InvoiceRepository interface {
	Create(ctx context.Context, c *models.Invoice) error
	GetByID(ctx context.Context, id string) (*models.Invoice, error)
	List(ctx context.Context) ([]*models.Invoice, error)
	Update(ctx context.Context, c *models.Invoice) error
	Delete(ctx context.Context, id string) error
}

type PostgresInvoiceRepository struct {
	Db *sql.DB
}

// Create implements InvoiceRepository
func (r *PostgresInvoiceRepository) Create(ctx context.Context, c *models.Invoice) error {
	query := `
		INSERT INTO invoices (id, cohort_id, client_name, egg_quantity, amount, status, due_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.Db.ExecContext(ctx, query,
		c.ID, c.CohortID, c.ClientName, c.EggQuantity, c.Amount, c.Status,
		c.DueDate, time.Now(), time.Now(),
	)
	return err
}

// GetByID implements InvoiceRepository
func (r *PostgresInvoiceRepository) GetByID(ctx context.Context, id string) (*models.Invoice, error) {
	query := `
		SELECT id, cohort_id, client_name, egg_quantity, amount, status, due_date, created_at, updated_at
		FROM invoices
		WHERE id = $1
	`
	row := r.Db.QueryRowContext(ctx, query, id)

	var inv models.Invoice
	err := row.Scan(
		&inv.ID, &inv.CohortID, &inv.ClientName, &inv.EggQuantity,
		&inv.Amount, &inv.Status, &inv.DueDate, &inv.CreatedAt, &inv.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &inv, err
}

// List implements InvoiceRepository
func (r *PostgresInvoiceRepository) List(ctx context.Context) ([]*models.Invoice, error) {
	query := `
		SELECT id, cohort_id, client_name, egg_quantity, amount, status, due_date, created_at, updated_at
		FROM invoices
	`
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []*models.Invoice
	for rows.Next() {
		var inv models.Invoice
		err := rows.Scan(
			&inv.ID, &inv.CohortID, &inv.ClientName, &inv.EggQuantity,
			&inv.Amount, &inv.Status, &inv.DueDate, &inv.CreatedAt, &inv.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, &inv)
	}

	return invoices, nil
}

// Update implements InvoiceRepository
func (r *PostgresInvoiceRepository) Update(ctx context.Context, c *models.Invoice) error {
	query := `
		UPDATE invoices
		SET cohort_id = $1, client_name = $2, egg_quantity = $3, amount = $4,
		    status = $5, due_date = $6, updated_at = $7
		WHERE id = $8
	`
	_, err := r.Db.ExecContext(ctx, query,
		c.CohortID, c.ClientName, c.EggQuantity, c.Amount, c.Status, c.DueDate, time.Now(), c.ID,
	)
	return err
}

// Delete implements InvoiceRepository
func (r *PostgresInvoiceRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM invoices WHERE id = $1`
	_, err := r.Db.ExecContext(ctx, query, id)
	return err
}

// Constructor
func NewInvoiceRepository(db *sql.DB) InvoiceRepository {
	return &PostgresInvoiceRepository{Db: db}
}
