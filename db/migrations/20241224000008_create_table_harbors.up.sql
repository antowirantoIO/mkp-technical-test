-- Create harbors table
CREATE TABLE harbors (
    id VARCHAR(36) NOT NULL,
    harbor_code VARCHAR(20) NOT NULL UNIQUE,
    harbor_name VARCHAR(255) NOT NULL,
    un_locode VARCHAR(10) NOT NULL UNIQUE,
    country VARCHAR(100) NOT NULL,
    province VARCHAR(100) NOT NULL,
    city VARCHAR(100) NOT NULL,
    address TEXT NOT NULL,
    postal_code VARCHAR(20) NOT NULL,
    latitude DECIMAL(10,8) NOT NULL,
    longitude DECIMAL(11,8) NOT NULL,
    harbor_type VARCHAR(100) NOT NULL,
    harbor_category VARCHAR(100) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'active',
    max_ship_length DECIMAL(10,2) NULL,
    max_ship_beam DECIMAL(10,2) NULL,
    max_ship_draft DECIMAL(10,2) NULL,
    max_ship_dwt DECIMAL(12,2) NULL,
    berth_count INT NOT NULL DEFAULT 0,
    crane_count INT NOT NULL DEFAULT 0,
    storage_capacity DECIMAL(15,2) NULL,
    water_depth DECIMAL(8,2) NOT NULL,
    tidal_range DECIMAL(8,2) NULL,
    working_hours VARCHAR(100) NOT NULL,
    timezone VARCHAR(50) NOT NULL,
    contact_person VARCHAR(255) NOT NULL,
    contact_phone VARCHAR(20) NOT NULL,
    contact_email VARCHAR(255) NOT NULL,
    website VARCHAR(500) NULL,
    has_customs BOOLEAN NOT NULL DEFAULT FALSE,
    has_quarantine BOOLEAN NOT NULL DEFAULT FALSE,
    has_pilotage BOOLEAN NOT NULL DEFAULT FALSE,
    has_tug_service BOOLEAN NOT NULL DEFAULT FALSE,
    has_bunkering BOOLEAN NOT NULL DEFAULT FALSE,
    has_repair_service BOOLEAN NOT NULL DEFAULT FALSE,
    has_waste BOOLEAN NOT NULL DEFAULT FALSE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    established_at BIGINT NULL,
    notes TEXT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    deleted_at BIGINT NULL,
    PRIMARY KEY (id)
);

-- Create indexes
CREATE INDEX idx_harbors_harbor_code ON harbors (harbor_code);
CREATE INDEX idx_harbors_un_locode ON harbors (un_locode);
CREATE INDEX idx_harbors_country ON harbors (country);
CREATE INDEX idx_harbors_province ON harbors (province);
CREATE INDEX idx_harbors_city ON harbors (city);
CREATE INDEX idx_harbors_harbor_type ON harbors (harbor_type);
CREATE INDEX idx_harbors_harbor_category ON harbors (harbor_category);
CREATE INDEX idx_harbors_status ON harbors (status);
CREATE INDEX idx_harbors_is_active ON harbors (is_active);
CREATE INDEX idx_harbors_deleted_at ON harbors (deleted_at);