-- Add refresh token fields to users table
ALTER TABLE users 
ADD COLUMN refresh_token VARCHAR(500) NULL,
ADD COLUMN refresh_expires_at BIGINT NULL;

-- Create index for refresh token
CREATE INDEX idx_users_refresh_token ON users (refresh_token);