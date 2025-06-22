# Gallus — Poultry Farm Management System

Gallus is a modern, modular and data-driven farm management system built to digitize poultry operations. It enables farm owners to efficiently track egg production, sales, feed consumption and expenditures across multiple bird cohorts with actionable insights.

## Features

### Sales Module
- Create and manage invoices for egg sales
- Support full or installment payments
- Auto-clear invoices when paid
- Generate monthly sales reports

### Production Module
- Record daily egg production by cohort
- Track totals by cohort and period (weekly, monthly, etc.)
- Analyze trends across time ranges

### Feeds Module
- Record feed and water intake per cohort
- Track feed purchases from suppliers
- Analyze consumption over time

### Expenditures Module
- Categorize and record expenses by purpose, amount, and cohort
- Monitor spending by category, cohort, and time range

## Tech Stack

| Layer       | Stack                                  |
|-------------|----------------------------------------|
| Backend     | Go (Standard Library only) + PostgreSQL |
| Frontend    | Reactnative + Tailwind CSS + shadcn/ui + Lucide Icons |
| Database    | PostgreSQL                             |
| Container   | Docker + Docker Compose (for local dev) |
| Versioning  | Git                                     |

## Project Structure

```
gallus/
├── backend/                    # Go backend (modular)
│   ├── cmd/                   # Entrypoint (main.go)
│   ├── internal/              # Domain logic (sales, feeds, etc.)
│   ├── api/                   # HTTP handlers and routing
│   ├── pkg/                   # Shared helpers (DB, config)
│   ├── models/                # Shared Go structs
│   └── migrations/            # SQL schema files
│
├── frontend/                   # React frontend
│   ├── public/                # Static files
│   ├── src/                   # Source code
│   │   ├── screens/           # Pages (Dashboard, Sales, etc.)
│   │   ├── components/        # UI components
│   │   ├── lib/               # Hooks, API utils
│   │   └── App.tsx            # App root
│
├── designs/                    # UI mockups and screen images
│
├── docker-compose.yml          # Local development stack
├── .env                        # Environment variables
└── README.md
```

## Getting Started

### Prerequisites
- [Go 1.21+](https://go.dev/)
- [PostgreSQL 14+](https://www.postgresql.org/)
- [Docker](https://www.docker.com/) (optional but recommended)

### Setup Instructions

#### 1. Clone the repository
```bash
git clone https://github.com/lenardjombo/gallus.git
cd gallus
```

#### 2. Setup environment variables
Create a `.env` file in the root directory:
```env
DB_URL=postgres://gallus:secret@localhost:5432/gallus_db?sslmode=disable
```

#### 3. Run with Docker Compose (recommended)
```bash
docker-compose up --build
```

This will:
- Start PostgreSQL
- Start your Go backend (on port 8080)

#### 4. Run the frontend
In another terminal:
```bash
cd frontend
npm install
npm run dev
```

Visit http://localhost:5173

## Reporting and Insights

Each module supports detailed, time-based analytics:
- Monthly sales reports
- Weekly and quarterly egg production trends
- Feed efficiency per cohort
- Expense breakdowns per category

## Testing

Coming soon: unit tests for services and HTTP handlers.

## Deployment

- Backend can be containerized and deployed to any cloud VM or Docker host
- PostgreSQL backups recommended via pg_dump
- Frontend can be exported and served via Nginx or Vercel

## UI Design

All mockup screens can be found in `/designs`. Each image represents a module layout and flow used to implement the frontend screens.

## Author

**Lenard Jombo**
- Computer Scientist  | Software Developer
- Kenya
- Passionate about Go, data-driven systems and modern backend architecture.

## License

MIT License — use freely, improve collaboratively.
