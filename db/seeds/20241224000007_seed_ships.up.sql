-- Seed data for ships table
-- Ships for maritime boarding system

INSERT INTO ships (id, ship_name, imo_number, call_sign, flag_state, ship_type, gross_tonnage, operator_id, created_at, updated_at) VALUES
('990e8400-e29b-41d4-a716-446655440001', 'MV Pacific Explorer', '9876543210', 'ABCD123', 'Panama', 'Container Ship', 75000, '880e8400-e29b-41d4-a716-446655440001', 1735027200, 1735027200),
('990e8400-e29b-41d4-a716-446655440002', 'SS Atlantic Voyager', '9876543211', 'EFGH456', 'Liberia', 'Bulk Carrier', 85000, '880e8400-e29b-41d4-a716-446655440002', 1735027200, 1735027200),
('990e8400-e29b-41d4-a716-446655440003', 'MV Global Trader', '9876543212', 'IJKL789', 'Marshall Islands', 'General Cargo', 45000, '880e8400-e29b-41d4-a716-446655440003', 1735027200, 1735027200),
('990e8400-e29b-41d4-a716-446655440004', 'MS Coastal Express', '9876543213', 'MNOP012', 'United States', 'Passenger Ferry', 12000, '880e8400-e29b-41d4-a716-446655440004', 1735027200, 1735027200),
('990e8400-e29b-41d4-a716-446655440005', 'MV Marine Carrier', '9876543214', 'QRST345', 'Singapore', 'Tanker', 95000, '880e8400-e29b-41d4-a716-446655440005', 1735027200, 1735027200);