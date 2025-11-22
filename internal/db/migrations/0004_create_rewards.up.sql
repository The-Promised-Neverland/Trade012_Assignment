CREATE TABLE IF NOT EXISTS rewards (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
    symbol VARCHAR(10) REFERENCES stocks(symbol),
    quantity NUMERIC(18,6) NOT NULL,
    buy_price NUMERIC(18,4) NOT NULL,
    reward_timestamp TIMESTAMP NOT NULL,
    UNIQUE(user_id, symbol, reward_timestamp)
);
