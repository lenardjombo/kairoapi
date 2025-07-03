package feeds

import (
	"context"
	"database/sql"
	"time"

	"github.com/lenardjombo/gallus/models"
)

type FeedsRepository interface {
	// Feed Purchases
	CreateFeedsPurchase(ctx context.Context, c *models.FeedPurchase) error
	GetFeedsPurchaseByID(ctx context.Context, id string) (*models.FeedPurchase, error)
	ListFeedsPurchases(ctx context.Context) ([]*models.FeedPurchase, error)
	UpdateFeedsPurchase(ctx context.Context, c *models.FeedPurchase) error
	DeleteFeedsPurchase(ctx context.Context, id string) error

	// Feed Consumption
	CreateFeedConsumption(ctx context.Context, c *models.FeedConsumption) error
	GetFeedConsumptionByID(ctx context.Context, id string) (*models.FeedConsumption, error)
	ListFeedConsumption(ctx context.Context) ([]*models.FeedConsumption, error)
	UpdateFeedConsumption(ctx context.Context, c *models.FeedConsumption) error
	DeleteFeedConsumption(ctx context.Context, id string) error
}


type PostgresFeedsRepository struct {
	DB *sql.DB
}



//create feeds purchase
func (r *PostgresFeedsRepository) CreateFeedsPurchase(ctx context.Context, c *models.FeedPurchase) error {
	query := `
		INSERT INTO feed_purchases (id, supplier_id, purchase_date, cost, bags, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.DB.ExecContext(ctx, query,
		c.ID, c.SupplierID, time.Now(), c.Cost, c.Bags, time.Now(),
	)
	return err
}

//Get a feeds purchase by id 
func (r *PostgresFeedsRepository) GetFeedsPurchaseByID(ctx context.Context, id string) (*models.FeedPurchase, error) {
	query := `SELECT id, supplier_id, purchase_date, cost, bags, created_at FROM feed_purchases WHERE id = $1`
	row := r.DB.QueryRowContext(ctx, query, id)

	var f models.FeedPurchase
	err := row.Scan(&f.ID, &f.SupplierID, &f.PurchaseDate, &f.Cost, &f.Bags, &f.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &f, nil
}

//List Feeds purchases
func (r *PostgresFeedsRepository) ListFeedsPurchases(ctx context.Context) ([]*models.FeedPurchase, error) {
	query := `SELECT id, supplier_id, purchase_date, cost, bags, created_at FROM feed_purchases ORDER BY purchase_date DESC`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var purchases []*models.FeedPurchase
	for rows.Next() {
		var f models.FeedPurchase
		if err := rows.Scan(&f.ID, &f.SupplierID, &f.PurchaseDate, &f.Cost, &f.Bags, &f.CreatedAt); err != nil {
			return nil, err
		}
		purchases = append(purchases, &f)
	}
	return purchases, nil
}

//Update Feeds purchases 
func (r *PostgresFeedsRepository) UpdateFeedsPurchase(ctx context.Context, f *models.FeedPurchase) error {
	query := `
		UPDATE feed_purchases
		SET supplier_id = $1, purchase_date = $2, cost = $3, bags = $4
		WHERE id = $5
	`
	_, err := r.DB.ExecContext(ctx, query, f.SupplierID, f.PurchaseDate, f.Cost, f.Bags, f.ID)
	return err
}

//Delete Feeds purchases 
func (r *PostgresFeedsRepository) DeleteFeedsPurchase(ctx context.Context, id string) error {
	query := `DELETE FROM feed_purchases WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}


//Create Feeds consumption
func (r *PostgresFeedsRepository) CreateFeedConsumption(ctx context.Context, c *models.FeedConsumption) error {
	query := `
		INSERT INTO feed_consumption (id, cohort_id, date, feed_kg, water_liters, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.DB.ExecContext(ctx, query,
		c.ID, c.CohortID, c.Date, c.FeedKG, c.WaterLiters, time.Now(),
	)
	return err
}

//Get Feeds Consumption by Id
func (r *PostgresFeedsRepository) GetFeedConsumptionByID(ctx context.Context, id string) (*models.FeedConsumption, error) {
	query := `SELECT id, cohort_id, date, feed_kg, water_liters, created_at FROM feed_consumption WHERE id = $1`
	row := r.DB.QueryRowContext(ctx, query, id)

	var f models.FeedConsumption
	err := row.Scan(&f.ID, &f.CohortID, &f.Date, &f.FeedKG, &f.WaterLiters, &f.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &f, nil
}


//Get List of all feeds consumption 
func (r *PostgresFeedsRepository) ListFeedConsumption(ctx context.Context) ([]*models.FeedConsumption, error) {
	query := `SELECT id, cohort_id, date, feed_kg, water_liters, created_at FROM feed_consumption ORDER BY date DESC`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var consumptions []*models.FeedConsumption
	for rows.Next() {
		var f models.FeedConsumption
		if err := rows.Scan(&f.ID, &f.CohortID, &f.Date, &f.FeedKG, &f.WaterLiters, &f.CreatedAt); err != nil {
			return nil, err
		}
		consumptions = append(consumptions, &f)
	}
	return consumptions, nil
}


//Update Feeds consumption
func (r *PostgresFeedsRepository) UpdateFeedConsumption(ctx context.Context, f *models.FeedConsumption) error {
	query := `
		UPDATE feed_consumption
		SET cohort_id = $1, date = $2, feed_kg = $3, water_liters = $4
		WHERE id = $5
	`
	_, err := r.DB.ExecContext(ctx, query, f.CohortID, f.Date, f.FeedKG, f.WaterLiters, f.ID)
	return err
}


//Delete Feeds consumption
func (r *PostgresFeedsRepository) DeleteFeedConsumption(ctx context.Context, id string) error {
	query := `DELETE FROM feed_consumption WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}

func NewFeedsRepository(database *sql.DB) FeedsRepository {
	return &PostgresFeedsRepository{DB: database}
}
