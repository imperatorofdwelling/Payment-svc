CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE PAYMENT_STATUS AS ENUM ('pending', 'waiting_for_capture', 'succeeded', 'canceled');
CREATE TYPE PAYMENT_AMOUNT_CURRENCY AS ENUM ('RUB', 'USD');

CREATE TABLE IF NOT EXISTS payments_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    transaction_id UUID NOT NULL UNIQUE,
    status PAYMENT_STATUS NOT NULL DEFAULT 'pending',
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS payments_logs_amount (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    value numeric(10,2) NOT NULL,
    currency PAYMENT_AMOUNT_CURRENCY NOT NULL,
    payment_id UUID UNIQUE NOT NULL REFERENCES payments_logs(id) ON DELETE CASCADE,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS payments_logs_method (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    payment_id UUID UNIQUE NOT NULL REFERENCES payments_logs(id) ON DELETE CASCADE,
    type varchar(255) NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);