-- Seed data for permissions table
-- Maritime boarding system permissions

INSERT INTO permissions (id, name, display_name, description, resource, action, is_active, is_system, created_at, updated_at, deleted_at) VALUES
-- User management permissions
('660e8400-e29b-41d4-a716-446655440001', 'user.index', 'View Users', 'View and list user accounts', 'user', 'index', true, true, 1735027200, 1735027200, NULL),
('660e8400-e29b-41d4-a716-446655440002', 'user.store', 'Create Users', 'Create new user accounts', 'user', 'store', true, true, 1735027200, 1735027200, NULL),
('660e8400-e29b-41d4-a716-446655440003', 'user.update', 'Update Users', 'Update existing user accounts', 'user', 'update', true, true, 1735027200, 1735027200, NULL),
('660e8400-e29b-41d4-a716-446655440004', 'user.destroy', 'Delete Users', 'Delete user accounts', 'user', 'destroy', true, true, 1735027200, 1735027200, NULL),
-- Harbor management permissions
('660e8400-e29b-41d4-a716-446655440005', 'harbor.index', 'View Harbors', 'View and list harbor information', 'harbor', 'index', true, true, 1735027200, 1735027200, NULL),
('660e8400-e29b-41d4-a716-446655440006', 'harbor.store', 'Create Harbors', 'Create new harbor records', 'harbor', 'store', true, true, 1735027200, 1735027200, NULL),
('660e8400-e29b-41d4-a716-446655440007', 'harbor.update', 'Update Harbors', 'Update existing harbor information', 'harbor', 'update', true, true, 1735027200, 1735027200, NULL),
('660e8400-e29b-41d4-a716-446655440008', 'harbor.destroy', 'Delete Harbors', 'Delete harbor records', 'harbor', 'destroy', true, true, 1735027200, 1735027200, NULL),
-- Ship management permissions
('660e8400-e29b-41d4-a716-446655440009', 'ship.index', 'View Ships', 'View and list ship records', 'ship', 'index', true, true, 1735027200, 1735027200, NULL),
('660e8400-e29b-41d4-a716-446655440010', 'ship.store', 'Create Ships', 'Create new ship records', 'ship', 'store', true, true, 1735027200, 1735027200, NULL),
('660e8400-e29b-41d4-a716-446655440011', 'ship.update', 'Update Ships', 'Update existing ship records', 'ship', 'update', true, true, 1735027200, 1735027200, NULL),
('660e8400-e29b-41d4-a716-446655440012', 'ship.destroy', 'Delete Ships', 'Delete ship records', 'ship', 'destroy', true, true, 1735027200, 1735027200, NULL),
-- Operator management permissions
('660e8400-e29b-41d4-a716-446655440013', 'operator.index', 'View Operators', 'View and list maritime operators', 'operator', 'index', true, true, 1735027200, 1735027200, NULL),
('660e8400-e29b-41d4-a716-446655440014', 'operator.store', 'Create Operators', 'Create new maritime operators', 'operator', 'store', true, true, 1735027200, 1735027200, NULL),
('660e8400-e29b-41d4-a716-446655440015', 'operator.update', 'Update Operators', 'Update existing maritime operators', 'operator', 'update', true, true, 1735027200, 1735027200, NULL),
('660e8400-e29b-41d4-a716-446655440016', 'operator.destroy', 'Delete Operators', 'Delete maritime operators', 'operator', 'destroy', true, true, 1735027200, 1735027200, NULL),
-- Boarding operations permissions
('660e8400-e29b-41d4-a716-446655440017', 'boarding.conduct', 'Conduct Boarding', 'Perform ship boarding operations and inspections', 'boarding', 'conduct', true, true, 1735027200, 1735027200, NULL),
-- Role management permissions
('660e8400-e29b-41d4-a716-446655440018', 'role.index', 'View Roles', 'View and list roles', 'role', 'index', true, true, 1735027200, 1735027200, NULL),
('660e8400-e29b-41d4-a716-446655440019', 'role.store', 'Create Roles', 'Create new roles', 'role', 'store', true, true, 1735027200, 1735027200, NULL),
('660e8400-e29b-41d4-a716-446655440020', 'role.update', 'Update Roles', 'Update existing roles', 'role', 'update', true, true, 1735027200, 1735027200, NULL),
('660e8400-e29b-41d4-a716-446655440021', 'role.destroy', 'Delete Roles', 'Delete roles', 'role', 'destroy', true, true, 1735027200, 1735027200, NULL),
-- Permission management permissions
('660e8400-e29b-41d4-a716-446655440022', 'permission.index', 'View Permissions', 'View and list permissions', 'permission', 'index', true, true, 1735027200, 1735027200, NULL),
('660e8400-e29b-41d4-a716-446655440023', 'permission.store', 'Create Permissions', 'Create new permissions', 'permission', 'store', true, true, 1735027200, 1735027200, NULL),
('660e8400-e29b-41d4-a716-446655440024', 'permission.update', 'Update Permissions', 'Update existing permissions', 'permission', 'update', true, true, 1735027200, 1735027200, NULL),
('660e8400-e29b-41d4-a716-446655440025', 'permission.destroy', 'Delete Permissions', 'Delete permissions', 'permission', 'destroy', true, true, 1735027200, 1735027200, NULL);