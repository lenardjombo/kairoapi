package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/lenardjombo/gallus/models"
)

type FeedPurchaseRepository interface {
	Create(ctx context.Context, fp *models.FeedPurchase) error
	GetByID(ctx context.Context, id string) (*models.FeedPurchase, error)
	List(ctx context.Context) ([]*models.FeedPurchase, error)
	Update(ctx context.Context, fp *models.FeedPurchase) error
	Delete(ctx context.Context, id string) error
}

type PostgresFeedPurchaseRepository struct {
	Db *sql.DB
}

// Create implements FeedPurchaseRepository
func (r *PostgresFeedPurchaseRepository) Create(ctx context.Context, fp *models.FeedPurchase) error {
	query := `
		INSERT INTO feed_purchases (id, supplier_id, purchase_date, cost, bags, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.Db.ExecContext(ctx, query,
		fp.ID, fp.SupplierID, fp.PurchaseDate, fp.Cost, fp.Bags, time.Now(),
	)
	return err
}

// GetByID implements FeedPurchaseRepository
func (r *PostgresFeedPurchaseRepository) GetByID(ctx context.Context, id string) (*models.FeedPurchase, error) {
	query := `
		SELECT id, supplier_id, purchase_date, cost, bags, created_at
		FROM feed_purchases
		WHERE id = $1
	`
	row := r.Db.QueryRowContext(ctx, query, id)

	var fp models.FeedPurchase
	err := row.Scan(
		&fp.ID, &fp.SupplierID, &fp.PurchaseDate, &fp.Cost, &fp.Bags, &fp.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &fp, err
}

// List implements FeedPurchaseRepository
func (r *PostgresFeedPurchaseRepository) List(ctx context.Context) ([]*models.FeedPurchase, error) {
	query := `
		SELECT id, supplier_id, purchase_date, cost, bags, created_at
		FROM feed_purchases
	`
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var purchases []*models.FeedPurchase
	for rows.Next() {
		var fp models.FeedPurchase
		err := rows.Scan(
			&fp.ID, &fp.SupplierID, &fp.PurchaseDate, &fp.Cost, &fp.Bags, &fp.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, &fp)
	}
	return purchases, nil
}

// Update implements FeedPurchaseRepository
func (r *PostgresFeedPurchaseRepository) Update(ctx context.Context, fp *models.FeedPurchase) error {
	query := `
		UPDATE feed_purchases
		SET supplier_id = $1, purchase_date = $2, cost = $3, bags = $4
		WHERE id = $5
	`
	_, err := r.Db.ExecContext(ctx, query,
		fp.SupplierID, fp.PurchaseDate, fp.Cost, fp.Bags, fp.ID,
	)
	return err
}

// Delete implements FeedPurchaseRepository
func (r *PostgresFeedPurchaseRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM feed_purchases WHERE id = $1`
	_, err := r.Db.ExecContext(ctx, query, id)
	return err
}

// Constructor
func NewFeedPurchaseRepository(db *sql.DB) FeedPurchaseRepository {
	return &PostgresFeedPurchaseRepository{Db: db}
}
