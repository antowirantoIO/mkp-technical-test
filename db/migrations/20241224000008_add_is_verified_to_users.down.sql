-- Remove is_verified column from users table
DROP INDEX IF EXISTS idx_users_is_verified;

ALTER TABLE users 
DROP COLUMN IF EXISTS is_verified;