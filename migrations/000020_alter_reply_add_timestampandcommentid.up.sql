ALTER TABLE IF EXISTS reply ADD COLUMN IF NOT EXISTS created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE IF EXISTS reply ADD COLUMN IF NOT EXISTS comment_id UUID REFERENCES comments(id) ON UPDATE CASCADE ON DELETE CASCADE;