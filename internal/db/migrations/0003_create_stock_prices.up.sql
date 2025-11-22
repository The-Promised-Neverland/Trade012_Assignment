CREATE TABLE stock_prices (
    symbol VARCHAR(10) PRIMARY KEY REFERENCES stocks(symbol),
    price NUMERIC(18,4) NOT NULL,
    price_timestamp TIMESTAMP NOT NULL
);