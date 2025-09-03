-- Create ships table
CREATE TABLE ships (
    id VARCHAR(36) PRIMARY KEY,
    operator_id VARCHAR(36) NOT NULL,
    ship_name VARCHAR(255) NOT NULL,
    imo_number VARCHAR(20) UNIQUE NOT NULL,
    call_sign VARCHAR(20) UNIQUE NOT NULL,
    mmsi VARCHAR(20) UNIQUE,
    ship_type VARCHAR(100) NOT NULL,
    flag_state VARCHAR(100) NOT NULL,
    port_of_registry VARCHAR(255),
    build_year INT,
    builder VARCHAR(255),
    length DECIMAL(10,2),
    beam DECIMAL(10,2),
    draft DECIMAL(10,2),
    gross_tonnage DECIMAL(12,2),
    net_tonnage DECIMAL(12,2),
    deadweight_tonnage DECIMAL(12,2),
    max_speed DECIMAL(5,2),
    passenger_capacity INT,
    crew_capacity INT,
    classification_society VARCHAR(255),
    status VARCHAR(50) DEFAULT 'active',
    is_active BOOLEAN DEFAULT TRUE,
    last_inspection BIGINT,
    next_inspection BIGINT,
    insurance_expiry BIGINT,
    certificate_expiry BIGINT,
    current_latitude DECIMAL(10,8),
    current_longitude DECIMAL(11,8),
    last_position BIGINT,
    notes TEXT,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    deleted_at BIGINT,
    
    FOREIGN KEY (operator_id) REFERENCES operators(id) ON DELETE CASCADE
);

-- Create indexes for ships table
CREATE INDEX idx_ships_operator_id ON ships(operator_id);
CREATE INDEX idx_ships_imo_number ON ships(imo_number);
CREATE INDEX idx_ships_call_sign ON ships(call_sign);
CREATE INDEX idx_ships_mmsi ON ships(mmsi);
CREATE INDEX idx_ships_status ON ships(status);
CREATE INDEX idx_ships_is_active ON ships(is_active);
CREATE INDEX idx_ships_deleted_at ON ships(deleted_at);