-- Create operators table
CREATE TABLE operators (
    id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL UNIQUE,
    operator_code VARCHAR(50) NOT NULL UNIQUE,
    company_name VARCHAR(255) NOT NULL,
    license_number VARCHAR(100) NOT NULL UNIQUE,
    contact_person VARCHAR(255) NOT NULL,
    contact_phone VARCHAR(20) NOT NULL,
    contact_email VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    city VARCHAR(100) NOT NULL,
    province VARCHAR(100) NOT NULL,
    country VARCHAR(100) NOT NULL,
    postal_code VARCHAR(20) NOT NULL,
    website VARCHAR(500) NULL,
    operator_type VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'active',
    established_at BIGINT NULL,
    license_expiry BIGINT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    notes TEXT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    deleted_at BIGINT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_operators_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- Create indexes
CREATE INDEX idx_operators_user_id ON operators (user_id);
CREATE INDEX idx_operators_operator_code ON operators (operator_code);
CREATE INDEX idx_operators_license_number ON operators (license_number);
CREATE INDEX idx_operators_operator_type ON operators (operator_type);
CREATE INDEX idx_operators_status ON operators (status);
CREATE INDEX idx_operators_is_active ON operators (is_active);
CREATE INDEX idx_operators_deleted_at ON operators (deleted_at);