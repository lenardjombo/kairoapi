CREATE TABLE "cohorts" (
  "id" uuid PRIMARY KEY,
  "name" varchar(100),
  "breed" varchar(100),
  "start_date" date,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "invoices" (
  "id" uuid PRIMARY KEY,
  "cohort_id" uuid,
  "client_name" varchar(100),
  "egg_quantity" int,
  "amount" decimal(10,2),
  "status" varchar(20),
  "due_date" date,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "payments" (
  "id" uuid PRIMARY KEY,
  "invoice_id" uuid,
  "amount" decimal(10,2),
  "paid_at" timestamp,
  "created_at" timestamp
);

CREATE TABLE "production_records" (
  "id" uuid PRIMARY KEY,
  "cohort_id" uuid,
  "date" date,
  "egg_count" int,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "suppliers" (
  "id" uuid PRIMARY KEY,
  "name" varchar(100),
  "contact" text,
  "created_at" timestamp
);

CREATE TABLE "feed_purchases" (
  "id" uuid PRIMARY KEY,
  "supplier_id" uuid,
  "purchase_date" date,
  "cost" decimal(10,2),
  "bags" int,
  "created_at" timestamp
);

CREATE TABLE "feed_consumption" (
  "id" uuid PRIMARY KEY,
  "cohort_id" uuid,
  "date" date,
  "feed_kg" decimal(10,2),
  "water_liters" decimal(10,2),
  "created_at" timestamp
);

CREATE TABLE "categories" (
  "id" uuid PRIMARY KEY,
  "name" varchar(100),
  "created_at" timestamp
);

CREATE TABLE "expenditures" (
  "id" uuid PRIMARY KEY,
  "category_id" uuid,
  "cohort_id" uuid,
  "amount" decimal(10,2),
  "name" varchar(100),
  "purpose" text,
  "date" date,
  "created_at" timestamp
);

ALTER TABLE "invoices" ADD FOREIGN KEY ("cohort_id") REFERENCES "cohorts" ("id");

ALTER TABLE "payments" ADD FOREIGN KEY ("invoice_id") REFERENCES "invoices" ("id");

ALTER TABLE "production_records" ADD FOREIGN KEY ("cohort_id") REFERENCES "cohorts" ("id");

ALTER TABLE "feed_purchases" ADD FOREIGN KEY ("supplier_id") REFERENCES "suppliers" ("id");

ALTER TABLE "feed_consumption" ADD FOREIGN KEY ("cohort_id") REFERENCES "cohorts" ("id");

ALTER TABLE "expenditures" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "expenditures" ADD FOREIGN KEY ("cohort_id") REFERENCES "cohorts" ("id");

