-- Create permissions table
CREATE TABLE permissions (
    id VARCHAR(36) NOT NULL,
    name VARCHAR(100) NOT NULL UNIQUE,
    display_name VARCHAR(255) NOT NULL,
    description TEXT NULL,
    resource VARCHAR(100) NOT NULL,
    action VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    is_system BOOLEAN NOT NULL DEFAULT FALSE,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    deleted_at BIGINT NULL,
    PRIMARY KEY (id),
    CONSTRAINT uk_permissions_resource_action UNIQUE (resource, action)
);

-- Create indexes
CREATE INDEX idx_permissions_name ON permissions (name);
CREATE INDEX idx_permissions_resource ON permissions (resource);
CREATE INDEX idx_permissions_action ON permissions (action);
CREATE INDEX idx_permissions_is_active ON permissions (is_active);
CREATE INDEX idx_permissions_is_system ON permissions (is_system);
CREATE INDEX idx_permissions_deleted_at ON permissions (deleted_at);