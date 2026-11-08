CREATE TABLE IF NOT EXISTS listings (
                                        id SERIAL PRIMARY KEY,
                                        title TEXT NOT NULL,
                                        price INT NOT NULL
);