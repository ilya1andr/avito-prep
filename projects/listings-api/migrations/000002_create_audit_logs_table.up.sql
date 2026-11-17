CREATE TABLE IF NOT EXISTS audit_logs (
                                          id SERIAL PRIMARY KEY,
                                          action TEXT NOT NULL,
                                          listing_id INT NOT NULL,
                                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);