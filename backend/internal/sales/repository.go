package sales

import (
	"context"
	"database/sql"
	"time"

	"github.com/lenardjombo/gallus/models"
)

type SalesRepository interface {
	// Invoice CRUD
	CreateInvoice(ctx context.Context, c *models.Invoice) error
	GetInvoiceByID(ctx context.Context, id string) (*models.Invoice, error)
	ListInvoices(ctx context.Context) ([]*models.Invoice, error)
	UpdateInvoice(ctx context.Context, c *models.Invoice) error
	DeleteInvoice(ctx context.Context, id string) error

	// Payment CRUD
	CreatePayment(ctx context.Context, c *models.Payment) error
	GetPaymentByID(ctx context.Context, id string) (*models.Payment, error)
	ListPayments(ctx context.Context) ([]*models.Payment, error)
	DeletePayment(ctx context.Context, id string) error
}

type PostgresSalesRepository struct {
	DB *sql.DB
}

// ========== Invoice Methods ==========

// CreateInvoice inserts a new invoice
func (r *PostgresSalesRepository) CreateInvoice(ctx context.Context, c *models.Invoice) error {
	query := `
		INSERT INTO invoices (id, cohort_id, client_name, egg_quantity, amount, status, due_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	now := time.Now()
	_, err := r.DB.ExecContext(ctx, query,
		c.ID,
		c.CohortID,
		c.ClientName,
		c.EggQuantity,
		c.Amount,
		c.Status,
		c.DueDate,
		now,
		now,
	)
	return err
}

// GetInvoiceByID fetches a single invoice
func (r *PostgresSalesRepository) GetInvoiceByID(ctx context.Context, id string) (*models.Invoice, error) {
	query := `
		SELECT id, cohort_id, client_name, egg_quantity, amount, status, due_date, created_at, updated_at
		FROM invoices
		WHERE id = $1
	`
	row := r.DB.QueryRowContext(ctx, query, id)

	var inv models.Invoice
	err := row.Scan(
		&inv.ID,
		&inv.CohortID,
		&inv.ClientName,
		&inv.EggQuantity,
		&inv.Amount,
		&inv.Status,
		&inv.DueDate,
		&inv.CreatedAt,
		&inv.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &inv, nil
}

// ListInvoices returns all invoices
func (r *PostgresSalesRepository) ListInvoices(ctx context.Context) ([]*models.Invoice, error) {
	query := `
		SELECT id, cohort_id, client_name, egg_quantity, amount, status, due_date, created_at, updated_at
		FROM invoices
		ORDER BY due_date DESC
	`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []*models.Invoice
	for rows.Next() {
		var inv models.Invoice
		err := rows.Scan(
			&inv.ID,
			&inv.CohortID,
			&inv.ClientName,
			&inv.EggQuantity,
			&inv.Amount,
			&inv.Status,
			&inv.DueDate,
			&inv.CreatedAt,
			&inv.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, &inv)
	}
	return invoices, nil
}

// UpdateInvoice modifies an existing invoice
func (r *PostgresSalesRepository) UpdateInvoice(ctx context.Context, c *models.Invoice) error {
	query := `
		UPDATE invoices
		SET cohort_id = $1, client_name = $2, egg_quantity = $3, amount = $4, status = $5, due_date = $6, updated_at = $7
		WHERE id = $8
	`
	_, err := r.DB.ExecContext(ctx, query,
		c.CohortID,
		c.ClientName,
		c.EggQuantity,
		c.Amount,
		c.Status,
		c.DueDate,
		time.Now(),
		c.ID,
	)
	return err
}

// DeleteInvoice removes an invoice by ID
func (r *PostgresSalesRepository) DeleteInvoice(ctx context.Context, id string) error {
	query := `DELETE FROM invoices WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}

// ========== Payment Methods ==========

// CreatePayment inserts a new payment
func (r *PostgresSalesRepository) CreatePayment(ctx context.Context, c *models.Payment) error {
	query := `
		INSERT INTO payments (id, invoice_id, amount, paid_at, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	now := time.Now()
	_, err := r.DB.ExecContext(ctx, query,
		c.ID,
		c.InvoiceID,
		c.Amount,
		now,
		now,
	)
	return err
}

// GetPaymentByID fetches a payment by ID
func (r *PostgresSalesRepository) GetPaymentByID(ctx context.Context, id string) (*models.Payment, error) {
	query := `
		SELECT id, invoice_id, amount, paid_at, created_at
		FROM payments
		WHERE id = $1
	`
	row := r.DB.QueryRowContext(ctx, query, id)

	var p models.Payment
	err := row.Scan(
		&p.ID,
		&p.InvoiceID,
		&p.Amount,
		&p.PaidAt,
		&p.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// ListPayments returns all payments
func (r *PostgresSalesRepository) ListPayments(ctx context.Context) ([]*models.Payment, error) {
	query := `
		SELECT id, invoice_id, amount, paid_at, created_at
		FROM payments
		ORDER BY paid_at DESC
	`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*models.Payment
	for rows.Next() {
		var p models.Payment
		err := rows.Scan(
			&p.ID,
			&p.InvoiceID,
			&p.Amount,
			&p.PaidAt,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		payments = append(payments, &p)
	}
	return payments, nil
}

// DeletePayment removes a payment by ID
func (r *PostgresSalesRepository) DeletePayment(ctx context.Context, id string) error {
	query := `DELETE FROM payments WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}

// Constructor - injects the DB dependency
func NewSalesRepository(database *sql.DB) SalesRepository {
	return &PostgresSalesRepository{DB: database}
}
