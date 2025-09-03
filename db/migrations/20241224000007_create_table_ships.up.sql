-- Create ships table
CREATE TABLE ships (
    id VARCHAR(36) NOT NULL,
    operator_id VARCHAR(36) NOT NULL,
    ship_name VARCHAR(255) NOT NULL,
    imo_number VARCHAR(20) NOT NULL UNIQUE,
    call_sign VARCHAR(20) NOT NULL UNIQUE,
    mmsi VARCHAR(20) NOT NULL UNIQUE,
    ship_type VARCHAR(100) NOT NULL,
    flag_state VARCHAR(100) NOT NULL,
    port_of_registry VARCHAR(255) NOT NULL,
    build_year INT NOT NULL,
    builder VARCHAR(255) NOT NULL,
    length DECIMAL(10,2) NOT NULL,
    beam DECIMAL(10,2) NOT NULL,
    draft DECIMAL(10,2) NOT NULL,
    gross_tonnage DECIMAL(12,2) NOT NULL,
    net_tonnage DECIMAL(12,2) NOT NULL,
    deadweight_tonnage DECIMAL(12,2) NOT NULL,
    max_speed DECIMAL(8,2) NOT NULL,
    passenger_capacity INT NOT NULL DEFAULT 0,
    crew_capacity INT NOT NULL DEFAULT 0,
    classification_society VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'active',
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    last_inspection BIGINT NULL,
    next_inspection BIGINT NULL,
    insurance_expiry BIGINT NULL,
    certificate_expiry BIGINT NULL,
    current_latitude DECIMAL(10,8) NULL,
    current_longitude DECIMAL(11,8) NULL,
    last_position BIGINT NULL,
    notes TEXT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    deleted_at BIGINT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_ships_operator_id FOREIGN KEY (operator_id) REFERENCES operators (id) ON DELETE CASCADE
);

-- Create indexes
CREATE INDEX idx_ships_operator_id ON ships (operator_id);
CREATE INDEX idx_ships_imo_number ON ships (imo_number);
CREATE INDEX idx_ships_call_sign ON ships (call_sign);
CREATE INDEX idx_ships_mmsi ON ships (mmsi);
CREATE INDEX idx_ships_ship_type ON ships (ship_type);
CREATE INDEX idx_ships_status ON ships (status);
CREATE INDEX idx_ships_is_active ON ships (is_active);
CREATE INDEX idx_ships_deleted_at ON ships (deleted_at);