-- Seed data for user_roles table
-- Assign roles to users based on their positions in maritime boarding system

INSERT INTO user_roles (user_id, role_id, created_at) VALUES
-- Admin System -> Super Admin
('770e8400-e29b-41d4-a716-446655440001', '550e8400-e29b-41d4-a716-446655440001', 1735027200),

-- Captain John Smith -> Port Authority
('770e8400-e29b-41d4-a716-446655440002', '550e8400-e29b-41d4-a716-446655440002', 1735027200),

-- Officer Sarah Wilson -> Boarding Officer
('770e8400-e29b-41d4-a716-446655440003', '550e8400-e29b-41d4-a716-446655440003', 1735027200),

-- Captain Michael Brown -> Ship Captain
('770e8400-e29b-41d4-a716-446655440004', '550e8400-e29b-41d4-a716-446655440004', 1735027200),

-- Manager Lisa Davis -> Operator Manager
('770e8400-e29b-41d4-a716-446655440005', '550e8400-e29b-41d4-a716-446655440005', 1735027200);