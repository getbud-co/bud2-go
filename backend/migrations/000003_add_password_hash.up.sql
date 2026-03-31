-- Add password_hash column to users table
ALTER TABLE users ADD COLUMN password_hash TEXT NOT NULL DEFAULT '';

-- Add comment explaining the column
COMMENT ON COLUMN users.password_hash IS 'Bcrypt hash of user password';
