CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    balance DECIMAL(20, 2) NOT NULL DEFAULT 0
);

INSERT INTO users (id, balance)
VALUES (1, 20), (2, 1.4), (3, 20012.12),
       (4, 0), (5, 27.2)
ON CONFLICT (id) DO NOTHING;

SELECT setval('users_id_seq', (SELECT MAX(id) FROM users));