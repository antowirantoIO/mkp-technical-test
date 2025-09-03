-- Seed data for roles table
-- Maritime boarding system roles

INSERT INTO roles (id, name, display_name, description, created_at, updated_at) VALUES
('550e8400-e29b-41d4-a716-446655440001', 'super_admin', 'super_admin', 'Super Administrator with full system access', 1735027200, 1735027200),
('550e8400-e29b-41d4-a716-446655440002', 'port_authority', 'port_authority', 'Port Authority Officer with harbor management access', 1735027200, 1735027200),
('550e8400-e29b-41d4-a716-446655440003', 'boarding_officer', 'boarding_officer', 'Boarding Officer for ship inspections and operations', 1735027200, 1735027200),
('550e8400-e29b-41d4-a716-446655440004', 'ship_captain', 'ship_captain', 'Ship Captain with vessel management access', 1735027200, 1735027200),
('550e8400-e29b-41d4-a716-446655440005', 'operator_manager', 'operator_manager', 'Operator Manager for maritime operations oversight', 1735027200, 1735027200);