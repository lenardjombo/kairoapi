// Db struct definitions
package models

import "time"

type Cohort struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    Breed     string    `json:"breed"`
    StartDate time.Time `json:"start_date"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Invoice struct {
    ID         string    `json:"id"`
    CohortID   string    `json:"cohort_id"`
    ClientName string    `json:"client_name"`
    EggQuantity int      `json:"egg_quantity"`
    Amount     string    `json:"amount"` // decimal, stored as string
    Status     string    `json:"status"`
    DueDate    time.Time `json:"due_date"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}

type Payment struct {
    ID        string    `json:"id"`
    InvoiceID string    `json:"invoice_id"`
    Amount    string    `json:"amount"` // decimal
    PaidAt    time.Time `json:"paid_at"`
    CreatedAt time.Time `json:"created_at"`
}

type ProductionRecord struct {
    ID        string    `json:"id"`
    CohortID  string    `json:"cohort_id"`
    Date      time.Time `json:"date"`
    EggCount  int       `json:"egg_count"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Supplier struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    Contact   string    `json:"contact"`
    CreatedAt time.Time `json:"created_at"`
}

type FeedPurchase struct {
    ID           string    `json:"id"`
    SupplierID   string    `json:"supplier_id"`
    PurchaseDate time.Time `json:"purchase_date"`
    Cost         string    `json:"cost"` // decimal
    Bags         int       `json:"bags"`
    CreatedAt    time.Time `json:"created_at"`
}

type FeedConsumption struct {
    ID          string    `json:"id"`
    CohortID    string    `json:"cohort_id"`
    Date        time.Time `json:"date"`
    FeedKG      string    `json:"feed_kg"`       // decimal
    WaterLiters string    `json:"water_liters"`  // decimal
    CreatedAt   time.Time `json:"created_at"`
}

type Category struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at"`
}

type Expenditure struct {
    ID         string    `json:"id"`
    CategoryID string    `json:"category_id"`
    CohortID   string    `json:"cohort_id"`
    Amount     string    `json:"amount"` // decimal
    Name       string    `json:"name"`
    Purpose    string    `json:"purpose"`
    Date       time.Time `json:"date"`
    CreatedAt  time.Time `json:"created_at"`
}
