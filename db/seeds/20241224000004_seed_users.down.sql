-- Remove seed data for users and users_extended tables

-- Delete from users_extended first (foreign key constraint)
DELETE FROM users_extended WHERE user_id IN (
    '770e8400-e29b-41d4-a716-446655440001',
    '770e8400-e29b-41d4-a716-446655440002',
    '770e8400-e29b-41d4-a716-446655440003',
    '770e8400-e29b-41d4-a716-446655440004',
    '770e8400-e29b-41d4-a716-446655440005'
);

-- Delete from users
DELETE FROM users WHERE id IN (
    '770e8400-e29b-41d4-a716-446655440001',
    '770e8400-e29b-41d4-a716-446655440002',
    '770e8400-e29b-41d4-a716-446655440003',
    '770e8400-e29b-41d4-a716-446655440004',
    '770e8400-e29b-41d4-a716-446655440005'
);