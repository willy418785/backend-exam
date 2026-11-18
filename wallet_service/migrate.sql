BEGIN;

CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    balance numeric(20, 2) NOT NULL DEFAULT 10000.00,
    CONSTRAINT accounts_valid_balance CHECK (
        balance >= 0
    )
);

CREATE TABLE transfers (
    id SERIAL PRIMARY KEY,
    from_id INT REFERENCES accounts (id),
    to_id INT REFERENCES accounts (id),
    amount numeric(20, 2) NOT NULL,
    CONSTRAINT transfers_valid_amount CHECK (
        amount > 0
    ),
    CONSTRAINT transfers_distinct_accounts CHECK (from_id != to_id)
);

COMMIT;
