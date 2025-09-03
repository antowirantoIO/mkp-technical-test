-- Seed data for harbors table
-- Harbors for maritime boarding system

INSERT INTO harbors (id, name, code, country, city, latitude, longitude, contact_email, contact_phone, created_at, updated_at) VALUES
('aa0e8400-e29b-41d4-a716-446655440001', 'Port of Singapore', 'SGSIN', 'Singapore', 'Singapore', 1.2966, 103.7764, 'info@singaporeport.com', '+65-6325-2493', 1735027200, 1735027200),
('aa0e8400-e29b-41d4-a716-446655440002', 'Port of Rotterdam', 'NLRTM', 'Netherlands', 'Rotterdam', 51.9225, 4.4792, 'contact@portofrotterdam.com', '+31-10-252-1010', 1735027200, 1735027200),
('aa0e8400-e29b-41d4-a716-446655440003', 'Port of Los Angeles', 'USLAX', 'United States', 'Los Angeles', 33.7361, -118.2922, 'info@portla.org', '+1-310-732-3508', 1735027200, 1735027200),
('aa0e8400-e29b-41d4-a716-446655440004', 'Port of Hamburg', 'DEHAM', 'Germany', 'Hamburg', 53.5511, 9.9937, 'contact@hafen-hamburg.de', '+49-40-37709-0', 1735027200, 1735027200),
('aa0e8400-e29b-41d4-a716-446655440005', 'Port of Hong Kong', 'HKHKG', 'Hong Kong', 'Hong Kong', 22.3193, 114.1694, 'enquiry@mardep.gov.hk', '+852-2852-4423', 1735027200, 1735027200);