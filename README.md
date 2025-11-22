Stocky - Stock Reward System
A backend service that manages stock rewards for users. Users receive fractional shares of Indian stocks (Reliance, TCS, Infosys, HDFC, ICICIBANK) as incentives. The system tracks these rewards, calculates associated fees (brokerage, STT, GST), and maintains a double-entry ledger to monitor all financial transactions.
🎯 Project Overview
Stocky is a hypothetical company that rewards users with shares of Indian stocks. When a user receives a reward:

They receive the full stock units without deductions
Stocky pays all brokerage, taxes, and regulatory fees from its side
A complete audit trail is maintained in the ledger system
Portfolio valuations are updated hourly with real-time stock prices
Historical portfolio snapshots are created daily at midnight

🛠 Tech Stack

Language: Go 1.20+
Web Framework: Gin
Database: PostgreSQL
Logging: Logrus
Task Scheduling: robfig/cron
Database ORM: sqlx
Migrations: golang-migrate

📋 Prerequisites
Before you begin, ensure you have the following installed:

Go 1.20 or higher
PostgreSQL 12 or higher
Git

🚀 Getting Started
1. Clone the Repository
bashgit clone https://github.com/The-Promised-Neverland/Trade012_Assignment.git
cd Trade012_Assignment
2. Set Up Environment Variables
Create a .env file in the project root:
env# Server Configuration
PORT=8080

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=your_password
DB_NAME=assignment
3. Create PostgreSQL Database
bashcreatedb assignment
Or using psql:
sqlCREATE DATABASE assignment;
4. Install Dependencies
bashgo mod download
go mod tidy
5. Run the Application
bashgo run main.go
The server will start on http://localhost:8080 and automatically run database migrations.
📚 API Endpoints
All endpoints are prefixed with /stocky
User Management
Get User Details
GET /stocky/user/:userId
Response:
json{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "created_at": "2024-01-15T10:30:00Z"
}
Reward Management
Grant Stock Reward
POST /stocky/reward
Request Body:
json{
  "user_id": 1,
  "stock": "RELIANCE",
  "shares": 10.5
}
Response:
json{
  "reward_id": 42,
  "message": "reward granted"
}
Fee Breakdown:

Stock Cost: (price × shares)
Brokerage: 0.03% of stock cost
STT (Securities Transaction Tax): 0.1% of stock cost
GST: 18% of brokerage fee
Total Cost = Stock Cost + Brokerage + STT + GST

Get Today's Rewards
GET /stocky/today-stocks/:userId
Response:
json[
  {
    "id": 40,
    "user_id": 1,
    "symbol": "RELIANCE",
    "quantity": 5.5,
    "buy_price": 2500.50,
    "reward_timestamp": "2024-01-20T09:00:00Z"
  },
  {
    "id": 41,
    "user_id": 1,
    "symbol": "TCS",
    "quantity": 3.2,
    "buy_price": 3600.00,
    "reward_timestamp": "2024-01-20T14:30:00Z"
  }
]
Get Portfolio Statistics
GET /stocky/stats/:userId
Response:
json{
  "today_shares": {
    "RELIANCE": 15.5,
    "TCS": 3.2
  },
  "total_portfolio_inr": 85000.50
}
Get Complete Portfolio
GET /stocky/portfolio/:userId
Response:
json{
  "holdings": {
    "RELIANCE": {
      "quantity": 25.5,
      "latest_price": 2520.00,
      "value_inr": 64260.00
    },
    "TCS": {
      "quantity": 8.2,
      "latest_price": 3650.50,
      "value_inr": 29934.10
    }
  },
  "total_portfolio_inr": 94194.10
}
Get Historical Portfolio Records
GET /stocky/historical-inr/:userId
Response:
json{
  "2024-01-18": 45000.00,
  "2024-01-19": 67500.25,
  "2024-01-20": 85000.50
}
Ledger Management
Get User Ledger
GET /stocky/ledger/:userId
Response:
json{
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
🗄️ Database Schema
Users Table
Stores user information who receive stock rewards.
Rewards Table
Records each stock reward event with timestamp and quantity.
Ledger Table
Double-entry system tracking all financial aspects of each reward including fees and taxes.
Stock Prices Table
Maintains the latest price for each tracked stock (updated every 10 seconds).
User Portfolio History Table
Daily snapshots of each user's portfolio (updated at midnight).
Stocks Table
Reference table for all tracked stock symbols.
⏰ Scheduled Tasks
Price Updater Cron (Every 10 seconds)

Fetches latest prices for tracked stocks (RELIANCE, TCS, INFOSYS, HDFC, ICICIBANK)
Generates random price variations to simulate market movement
Updates the stock_prices table with latest valuations

Portfolio Cron (Daily at Midnight)

Iterates through all users
Calculates total holdings per stock symbol
Fetches latest prices and calculates INR valuation
Creates daily portfolio snapshots in user_portfolio_history

📝 Key Features
1. Fractional Share Support

Uses NUMERIC(18, 6) for precise share quantities
Supports partial shares from various reward events

2. Comprehensive Fee Tracking

Automatic fee calculation on reward creation
Double-entry ledger for audit trail
Breakdowns for: brokerage, STT, GST, and other fees

3. Portfolio Valuation

Real-time portfolio value calculation
Historical snapshot tracking
Daily portfolio evolution records

4. Data Integrity

Unique constraints on reward events to prevent duplicates
Foreign key relationships for referential integrity
Transaction support for atomic operations

5. Logging & Monitoring

Structured JSON logging using Logrus
Detailed error tracking with context
Performance metrics in logs

🛡️ Edge Cases Handled
Duplicate Prevention

UNIQUE(user_id, symbol, reward_timestamp) constraint on rewards table prevents duplicate reward entries

Price Unavailability

System logs warnings when stock prices are unavailable
Defaults to 0 INR valuation for missing prices
Continues operation without blocking other rewards

Concurrent Updates

Uses database-level ON CONFLICT clauses for safe concurrent price updates
Proper transaction handling in all operations

Data Precision

Uses high-precision decimals (NUMERIC(18, 4) for INR, NUMERIC(18, 6) for shares)
Prevents rounding errors in financial calculations

User Validation

Middleware validates user existence before processing requests
Graceful error responses for invalid users

📊 Postman Collection
A complete Postman collection is included in the repository (postman_collection.json) with all API endpoints, example requests, and responses.
Import to Postman:

Open Postman
Click "Import"
Select the postman_collection.json file
All endpoints will be ready to use

main.go               # Application entry point
🚨 Error Handling
The API returns standard HTTP status codes:

200 OK - Successful operation
400 Bad Request - Invalid input or user ID
404 Not Found - User not found
500 Internal Server Error - Server-side errors

All errors are logged with context for debugging.
📖 Database Migrations
Migrations are automatically applied on startup. They create all necessary tables and relationships. If you need to add custom migrations:

Create migration files in internal/db/migrations/
Follow the naming convention: 000X_description.up.sql and 000X_description.down.sql
Restart the application

🔄 Workflow Example

User is created in the system
Admin grants 10 RELIANCE shares via /stocky/reward endpoint
System records the reward and calculates all fees
Ledger entry is created with complete fee breakdown
Price updater cron fetches latest RELIANCE price
User checks portfolio via /stocky/portfolio/:userId
At midnight, portfolio snapshot is created for historical tracking
