-- Remove seed data for harbors table

DELETE FROM harbors WHERE id IN (
    'aa0e8400-e29b-41d4-a716-446655440001',
    'aa0e8400-e29b-41d4-a716-446655440002',
    'aa0e8400-e29b-41d4-a716-446655440003',
    'aa0e8400-e29b-41d4-a716-446655440004',
    'aa0e8400-e29b-41d4-a716-446655440005'
);