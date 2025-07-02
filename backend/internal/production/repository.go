// Data Access layer for Production
package production

import (
	"context"
	"database/sql"
	"time"

	"github.com/lenardjombo/gallus/models"
)

//Interface - Defines the methods for interacting with Production data in the database
type ProductionRepository interface{
	Create(ctx context.Context,c *models.ProductionRecord) error
}

// Provides methods to interact with the Production Recoed Database 
type ProductionRepo struct {
	DB *sql.DB
}

func (r *ProductionRepo) Create(ctx context.Context,c *models.ProductionRecord) error{
	query := `
		INSERT INTO production_records (id, cohort_id, date, egg_count, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_,err := r.DB.ExecContext(ctx, query, 
		c.ID, 
		c.CohortID, 
		c.Date, 
		c.EggCount, 
		time.Now(), //created_at
		time.Now(),//updated_at
	) 

	return err
}

// Constructor - An new instance of ProductionRepo (Database Ijection) with the provided database 
func NewProductionRepository(database *sql.DB) ProductionRepository{
	return &ProductionRepo{DB: database}
}
