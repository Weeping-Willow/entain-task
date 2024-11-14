DO $$ BEGIN
    CREATE TYPE transaction_state AS ENUM ('win', 'lose');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
    CREATE TYPE transaction_source_type AS ENUM ('game', 'server', 'payment');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE IF NOT EXISTS transactions (
    id VARCHAR(256) PRIMARY KEY,
    user_id BIGINT NOT NULL,
    amount DECIMAL(20, 2) NOT NULL,
    state transaction_state NOT NULL,
    source_type transaction_source_type NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);