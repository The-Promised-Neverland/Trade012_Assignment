# Stocky - Stock Reward System

A backend service that manages stock rewards for users. Users receive fractional shares of Indian stocks as incentives. The system tracks rewards, calculates fees (brokerage, STT, GST), and maintains a double-entry ledger.

## Project Overview

Stocky is a hypothetical company that rewards users with shares of Indian stocks (Reliance, TCS, Infosys, HDFC, ICICIBANK). When a user receives a reward:
- They receive full stock units without deductions
- Stocky pays all brokerage, taxes, and regulatory fees
- Complete audit trail maintained in ledger
- Portfolio valuations updated hourly with stock prices
- Historical snapshots created daily at midnight

## Tech Stack

- Go 1.20+
- Gin (Web Framework)
- PostgreSQL (Database)
- Logrus (Logging)
- robfig/cron (Task Scheduling)
- sqlx (Database ORM)
- golang-migrate (Migrations)

## Prerequisites

- Go 1.20 or higher
- PostgreSQL 12 or higher
- Docker & Docker Compose (optional)
- Git

## Getting Started

### 1. Clone Repository

```bash
git clone https://github.com/The-Promised-Neverland/Trade012_Assignment.git
cd Trade012_Assignment
```

### 2. Create .env File

Create `.env` in project root:

```
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=assignment
```

### 3. Start PostgreSQL with Docker

```bash
docker-compose up -d
```

Or create database manually:

```bash
createdb assignment
```

### 4. Install Dependencies

```bash
go mod download
go mod tidy
```

### 5. Run Application

```bash
go run main.go
```

Server runs on `http://localhost:8080`

## API Endpoints

All endpoints prefixed with `/stocky`

### Users

**Get User**
```
GET /stocky/user/:userId
```

Response:
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "created_at": "2024-01-15T10:30:00Z"
}
```

### Rewards

**Grant Stock Reward**
```
POST /stocky/reward
```

Request:
```json
{
  "user_id": 1,
  "stock": "RELIANCE",
  "shares": 10.5
}
```

Response:
```json
{
  "reward_id": 42,
  "message": "reward granted"
}
```

**Fee Calculation:**
- Stock Cost = price × shares
- Brokerage = 0.03% of stock cost
- STT = 0.1% of stock cost
- GST = 18% of brokerage
- Total = Stock Cost + Brokerage + STT + GST

**Get Today's Rewards**
```
GET /stocky/today-stocks/:userId
```

Response:
```json
[
  {
    "id": 40,
    "user_id": 1,
    "symbol": "RELIANCE",
    "quantity": 5.5,
    "buy_price": 2500.50,
    "reward_timestamp": "2024-01-20T09:00:00Z"
  }
]
```

**Get Portfolio Stats**
```
GET /stocky/stats/:userId
```

Response:
```json
{
  "today_shares": {
    "RELIANCE": 15.5,
    "TCS": 3.2
  },
  "total_portfolio_inr": 85000.50
}
```

**Get Complete Portfolio**
```
GET /stocky/portfolio/:userId
```

Response:
```json
{
  "holdings": {
    "RELIANCE": {
      "quantity": 25.5,
      "latest_price": 2520.00,
      "value_inr": 64260.00
    }
  },
  "total_portfolio_inr": 94194.10
}
```

**Get Historical Records**
```
GET /stocky/historical-inr/:userId
```

Response:
```json
{
  "2024-01-18": 45000.00,
  "2024-01-19": 67500.25,
  "2024-01-20": 85000.50
}
```

### Ledger

**Get User Ledger**
```
GET /stocky/ledger/:userId
```

Response:
```json
{
  "userId": 1,
  "ledger": [
    {
      "id": 1,
      "reward_id": 40,
      "symbol": "RELIANCE",
      "quantity": 5.5,
      "inr_cost": 13756.50,
      "brokerage_fee": 41.27,
      "stt_tax": 13.75,
      "gst_fee": 7.43,
      "other_fees": 0,
      "created_at": "2024-01-20T09:00:00Z"
    }
  ]
}
```

## Database Schema

**Users Table** - Stores user information

**Rewards Table** - Records each stock reward event

**Ledger Table** - Tracks all financial transactions and fees

**Stock Prices Table** - Latest prices for each stock (updated every 10 seconds)

**User Portfolio History** - Daily portfolio snapshots (updated at midnight)

**Stocks Table** - Reference table for stock symbols

## Scheduled Tasks

**Price Updater (Every 10 seconds)**
- Fetches latest prices for: RELIANCE, TCS, INFOSYS, HDFC, ICICIBANK
- Updates stock_prices table with new valuations

**Portfolio Cron (Daily at Midnight)**
- Iterates through all users
- Calculates total holdings per stock
- Creates daily portfolio snapshots

## Key Features

**Fractional Share Support**
- Uses NUMERIC(18, 6) for precise quantities
- Supports partial shares

**Comprehensive Fee Tracking**
- Automatic fee calculation
- Double-entry ledger for audit trail
- Breakdowns for all fee types

**Portfolio Valuation**
- Real-time portfolio value calculation
- Historical snapshot tracking
- Daily portfolio evolution

**Data Integrity**
- Unique constraints prevent duplicates
- Foreign key relationships maintained
- Transaction support for atomic operations

**Logging & Monitoring**
- Structured JSON logging
- Detailed error tracking
- Performance metrics

## Edge Cases Handled

**Duplicate Prevention**
- UNIQUE constraint on (user_id, symbol, reward_timestamp)

**Price Unavailability**
- Warns when prices unavailable
- Defaults to 0 INR valuation
- Continues operation

**Concurrent Updates**
- Database-level ON CONFLICT clauses
- Proper transaction handling

**Data Precision**
- High-precision decimals for financial data
- Prevents rounding errors

**User Validation**
- Middleware validates user existence
- Graceful error responses

## Using Postman Collection

Import `postman_collection.json` to Postman:

1. Open Postman
2. Click "Import"
3. Select `postman_collection.json`
4. All endpoints ready to use

## HTTP Status Codes

- 200 OK - Success
- 400 Bad Request - Invalid input
- 404 Not Found - User not found
- 500 Internal Server Error - Server error

## Workflow Example

1. User created in system
2. Admin grants 10 RELIANCE shares via POST /stocky/reward
3. System records reward and calculates fees
4. Ledger entry created with fee breakdown
5. Price updater fetches latest RELIANCE price
6. User checks portfolio via GET /stocky/portfolio/:userId
7. At midnight, portfolio snapshot created

## Environment Variables

```
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=assignment
```

## Docker Commands

```bash
# Start services
docker-compose up -d

# Stop services
docker-compose down

# View logs
docker-compose logs postgres

# View running containers
docker-compose ps
```

## Requirements Met

✅ REST APIs for all required endpoints
✅ Double-entry ledger system
✅ Fractional share support (NUMERIC 18,6)
✅ Precise INR amounts (NUMERIC 18,4)
✅ Duplicate prevention with unique constraints
✅ Price API downtime handling
✅ Rounding error prevention
✅ Golang with Gin and Logrus
✅ PostgreSQL database named "assignment"
✅ Postman collection included
✅ .env file support
