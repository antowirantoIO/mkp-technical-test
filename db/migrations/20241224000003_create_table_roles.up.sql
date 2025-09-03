-- Create roles table
CREATE TABLE roles (
    id VARCHAR(36) NOT NULL,
    name VARCHAR(100) NOT NULL UNIQUE,
    display_name VARCHAR(255) NOT NULL,
    description TEXT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    is_system BOOLEAN NOT NULL DEFAULT FALSE,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    deleted_at BIGINT NULL,
    PRIMARY KEY (id)
);

-- Create indexes
CREATE INDEX idx_roles_name ON roles (name);
CREATE INDEX idx_roles_is_active ON roles (is_active);
CREATE INDEX idx_roles_is_system ON roles (is_system);
CREATE INDEX idx_roles_deleted_at ON roles (deleted_at);