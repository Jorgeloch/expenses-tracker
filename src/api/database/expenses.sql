CREATE TABLE IF NOT EXISTS expenses (
  id SERIAL PRIMARY KEY,
  value DECIMAL(10, 2) NOT NULL,
  date DATE NOT NULL,
  user_id INTEGER NOT NULL,
  card_id INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (card_id) REFERENCES cards (id),
);
