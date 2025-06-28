# Gallus Database Schema Documentation

This document describes the PostgreSQL database schema used in **Gallus**, a poultry farm management system. The schema is modular and supports tracking of sales, production, feed intake and expenditures across bird cohorts.

---

##  Modules Overview

| Module          | Tables |
|------------------|--------|
| Shared Entities  | `cohorts` |
| Sales            | `invoices`, `payments` |
| Production       | `production_records` |
| Feeds            | `suppliers`, `feed_purchases`, `feed_consumption` |
| Expenditures     | `categories`, `expenditures` |

---

##  Shared Entity

### `cohorts`

Represents a group of birds (e.g., by age, breed or cycle).

| Column       | Type          | Description            |
|--------------|---------------|------------------------|
| `id`         | `uuid`        | Primary key            |
| `name`       | `varchar(100)`| Cohort name            |
| `breed`      | `varchar(100)`| Bird breed             |
| `start_date` | `date`        | When the cohort started|
| `created_at` | `timestamp`   | Record creation time   |
| `updated_at` | `timestamp`   | Record last updated    |

---

##  Sales Module

###  `invoices`

Invoices generated for egg sales.

| Column         | Type            | Description                         |
|----------------|------------------|-------------------------------------|
| `id`           | `uuid`           | Primary key                         |
| `cohort_id`    | `uuid`           | FK → `cohorts.id`                   |
| `client_name`  | `varchar(100)`   | Name of the client                  |
| `egg_quantity` | `int`            | Number of eggs sold                 |
| `amount`       | `decimal(10, 2)` | Total amount                        |
| `status`       | `varchar(20)`    | Status: `unpaid`, `partial`, `paid`|
| `due_date`     | `date`           | Payment due date                    |
| `created_at`   | `timestamp`      | Invoice creation time               |
| `updated_at`   | `timestamp`      | Last updated                        |

###  `payments`

Tracks payments made against invoices.

| Column       | Type            | Description                   |
|--------------|------------------|-------------------------------|
| `id`         | `uuid`           | Primary key                   |
| `invoice_id` | `uuid`           | FK → `invoices.id`            |
| `amount`     | `decimal(10, 2)` | Amount paid                   |
| `paid_at`    | `timestamp`      | Payment date                  |
| `created_at` | `timestamp`      | Record creation time          |

---

##  Production Module

###  `production_records`

Daily records of egg production per cohort.

| Column       | Type          | Description                |
|--------------|---------------|----------------------------|
| `id`         | `uuid`        | Primary key                |
| `cohort_id`  | `uuid`        | FK → `cohorts.id`          |
| `date`       | `date`        | Date of record             |
| `egg_count`  | `int`         | Number of eggs produced    |
| `created_at` | `timestamp`   | Record creation time       |
| `updated_at` | `timestamp`   | Last updated               |

---

##  Feeds Module

###  `suppliers`

Suppliers from whom feed is purchased.

| Column       | Type          | Description           |
|--------------|---------------|-----------------------|
| `id`         | `uuid`        | Primary key           |
| `name`       | `varchar(100)`| Supplier name         |
| `contact`    | `text`        | Contact info          |
| `created_at` | `timestamp`   | Record creation time  |

###  `feed_purchases`

Feed purchases linked to suppliers.

| Column         | Type             | Description                |
|----------------|------------------|----------------------------|
| `id`           | `uuid`           | Primary key                |
| `supplier_id`  | `uuid`           | FK → `suppliers.id`        |
| `purchase_date`| `date`           | Date of purchase           |
| `cost`         | `decimal(10, 2)` | Total cost of purchase     |
| `bags`         | `int`            | Number of feed bags bought |
| `created_at`   | `timestamp`      | Record creation time       |

###  `feed_consumption`

Daily feed and water intake per cohort.

| Column         | Type             | Description                    |
|----------------|------------------|--------------------------------|
| `id`           | `uuid`           | Primary key                    |
| `cohort_id`    | `uuid`           | FK → `cohorts.id`              |
| `date`         | `date`           | Record date                    |
| `feed_kg`      | `decimal(10, 2)` | Feed consumed in kilograms     |
| `water_liters` | `decimal(10, 2)` | Water consumed in liters       |
| `created_at`   | `timestamp`      | Record creation time           |

---

##  Expenditures Module
###  `categories`

Categories for labeling expenditures.

| Column       | Type          | Description           |
|--------------|---------------|-----------------------|
| `id`         | `uuid`        | Primary key           |
| `name`       | `varchar(100)`| Category name         |
| `created_at` | `timestamp`   | Record creation time  |

###  `expenditures`

Tracks spending across categories and cohorts.

| Column       | Type             | Description                        |
|--------------|------------------|------------------------------------|
| `id`         | `uuid`           | Primary key                        |
| `category_id`| `uuid`           | FK → `categories.id`               |
| `cohort_id`  | `uuid`           | FK → `cohorts.id`                  |
| `amount`     | `decimal(10, 2)` | Amount spent                       |
| `name`       | `varchar(100)`   | Short label for the expense        |
| `purpose`    | `text`           | Description or justification       |
| `date`       | `date`           | When the expense occurred          |
| `created_at` | `timestamp`      | Record creation time               |

---


Related Files

[gallus_schema.sql](""https://github.com/lenardjombo/gallus/blob/main/docs/database/gallusdb.sql) – SQL file to create the entire schema

[gallus_db_design]("https://github.com/lenardjombo/gallus/blob/main/docs/database/schema-design.pdf") – Visual ER diagram
