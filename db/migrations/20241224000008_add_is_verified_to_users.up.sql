-- Add is_verified column to users table
ALTER TABLE users 
ADD COLUMN is_verified BOOLEAN NOT NULL DEFAULT FALSE;

ALTER TABLE users ADD COLUMN IF NOT EXISTS password_changed_at BIGINT NULL;

-- Create index for is_verified
CREATE INDEX idx_users_is_verified ON users (is_verified);

-- Create index for password_changed_at
CREATE INDEX idx_users_password_changed_at ON users (password_changed_at);