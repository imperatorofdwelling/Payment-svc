CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    transaction_id varchar(255) NOT NULL UNIQUE,
    transaction_type varchar(255) NOT NULL,
    status varchar(255) NOT NULL DEFAULT 'pending',
    value numeric(10,2) NOT NULL,
    currency varchar(10) NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);