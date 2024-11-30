-- +goose Up
-- +goose StatementBegin
CREATE TYPE payment_status AS ENUM ('Success', 'In progress', 'Cancelled');
CREATE TABLE IF NOT EXISTS payment (
    id BIGSERIAL PRIMARY KEY,
    Sum DECIMAL(12, 2),
    payment_date TIMESTAMP,
    status payment_status,    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payment;
-- +goose StatementEnd
