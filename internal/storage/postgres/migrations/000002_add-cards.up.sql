CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS bank_cards (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    issuer_name varchar(255),
    issuer_country varchar(255),
    payout_token varchar(100) NOT NULL UNIQUE,
    first6 varchar(6) NOT NULL,
    last4 varchar(4) NOT NULL,
    card_type varchar(100) NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);