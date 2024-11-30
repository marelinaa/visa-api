-- +goose Up
-- +goose StatementBegin
CREATE TYPE application_status AS ENUM ('Approved', 'In progress', 'Denied');
CREATE TABLE IF NOT EXISTS application (
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    phone_number VARCHAR(20) NOT NULL,
    email VARCHAR(255),
    submission_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status application_status,
    applicant_id  BIGINT,
    operator_id BIGINT,
    payment_id BIGINT,
    FOREIGN KEY (applicant_id) REFERENCES applicant(id),
    FOREIGN KEY (operator_id) REFERENCES operator(id),
    FOREIGN KEY (payment_id) REFERENCES payment(id),
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS application;
-- +goose StatementEnd
