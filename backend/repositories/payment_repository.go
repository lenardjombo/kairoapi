package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/lenardjombo/gallus/models"
)

type PaymentRepository interface {
	Create(ctx context.Context, p *models.Payment) error
	GetByID(ctx context.Context, id string) (*models.Payment, error)
	List(ctx context.Context) ([]*models.Payment, error)
	Update(ctx context.Context, p *models.Payment) error
	Delete(ctx context.Context, id string) error
}

type PostgresPaymentRepository struct {
	Db *sql.DB
}

// Create implements PaymentRepository
func (r *PostgresPaymentRepository) Create(ctx context.Context, p *models.Payment) error {
	query := `
		INSERT INTO payments (id, invoice_id, amount, paid_at, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.Db.ExecContext(ctx, query,
		p.ID, p.InvoiceID, p.Amount, p.PaidAt, time.Now(),
	)
	return err
}

// GetByID implements PaymentRepository
func (r *PostgresPaymentRepository) GetByID(ctx context.Context, id string) (*models.Payment, error) {
	query := `
		SELECT id, invoice_id, amount, paid_at, created_at
		FROM payments
		WHERE id = $1
	`
	row := r.Db.QueryRowContext(ctx, query, id)

	var p models.Payment
	err := row.Scan(&p.ID, &p.InvoiceID, &p.Amount, &p.PaidAt, &p.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &p, err
}

// List implements PaymentRepository
func (r *PostgresPaymentRepository) List(ctx context.Context) ([]*models.Payment, error) {
	query := `
		SELECT id, invoice_id, amount, paid_at, created_at
		FROM payments
	`
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*models.Payment
	for rows.Next() {
		var p models.Payment
		err := rows.Scan(&p.ID, &p.InvoiceID, &p.Amount, &p.PaidAt, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		payments = append(payments, &p)
	}
	return payments, nil
}

// Update implements PaymentRepository
func (r *PostgresPaymentRepository) Update(ctx context.Context, p *models.Payment) error {
	query := `
		UPDATE payments
		SET invoice_id = $1, amount = $2, paid_at = $3
		WHERE id = $4
	`
	_, err := r.Db.ExecContext(ctx, query, p.InvoiceID, p.Amount, p.PaidAt, p.ID)
	return err
}

// Delete implements PaymentRepository
func (r *PostgresPaymentRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM payments WHERE id = $1`
	_, err := r.Db.ExecContext(ctx, query, id)
	return err
}

// Constructor
func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &PostgresPaymentRepository{Db: db}
}
