package sales

import(
	"context"
	"database/sql"
	"time"

	"github.com/lenardjombo/gallus/models"
)

type SalesRepository interface {
	CreatePayment(ctx context.Context,c *models.Payment)error
	CreateInvoice(ctx context.Context,c *models.Invoice)error
}

type PostgresSalesRepository struct {
	DB *sql.DB
}

//Create Payment
func (r *PostgresSalesRepository) CreatePayment(ctx context.Context,c *models.Payment)error{
	query := `
		INSERT INTO payments (id, invoice_id, amount, paid_at, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_,err := r.DB.ExecContext(ctx, query,
		 c.ID, 
		 c.InvoiceID,
		 c.Amount, 
		 time.Now(), //paid_at 
		 time.Now()) //created_at
	return err
}

//Create Invoice
func (r *PostgresSalesRepository)CreateInvoice(ctx context.Context, c *models.Invoice)error{
	query := `
		INSERT INTO invoices (id, cohort_id, client_name, egg_quantity, amount, status, due_date, create_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_,err := r.DB.ExecContext(ctx,query,
		c.ID, 
		c.CohortID,
		c.ClientName,
		c.EggQuantity,
		c.Amount,
		c.Status,
		time.Now(),// due_date
		time.Now(),// created_at
	)
	return err
}

func NewSalesRepository(database *sql.DB) SalesRepository{
	return &PostgresSalesRepository{DB: database}
}