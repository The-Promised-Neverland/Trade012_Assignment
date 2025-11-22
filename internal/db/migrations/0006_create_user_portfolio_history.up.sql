CREATE TABLE IF NOT EXISTS user_portfolio_history (
	user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
	date DATE NOT NULL,
	symbol VARCHAR(10) REFERENCES stocks(symbol),
	total_shares NUMERIC(18,6) NOT NULL,
	value_inr NUMERIC(18,4) NOT NULL,
	PRIMARY KEY(user_id, date, symbol)
);
