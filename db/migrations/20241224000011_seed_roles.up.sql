-- Seed data for roles table
-- Maritime boarding system roles

INSERT INTO roles (id, name, display_name, description, is_active, is_system, created_at, updated_at, deleted_at) VALUES
('550e8400-e29b-41d4-a716-446655440001', 'super_admin', 'Super Administrator', 'Super Administrator with full system access', true, true, 1735027200, 1735027200, NULL),
('550e8400-e29b-41d4-a716-446655440002', 'port_authority', 'Port Authority', 'Port Authority Officer with harbor management access', true, true, 1735027200, 1735027200, NULL),
('550e8400-e29b-41d4-a716-446655440003', 'boarding_officer', 'Boarding Officer', 'Boarding Officer for ship inspections and operations', true, true, 1735027200, 1735027200, NULL),
('550e8400-e29b-41d4-a716-446655440004', 'ship_captain', 'Ship Captain', 'Ship Captain with vessel management access', true, true, 1735027200, 1735027200, NULL),
('550e8400-e29b-41d4-a716-446655440005', 'operator_manager', 'Operator Manager', 'Operator Manager for maritime operations oversight', true, true, 1735027200, 1735027200, NULL);