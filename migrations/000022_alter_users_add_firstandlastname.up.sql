ALTER TABLE users ADD COLUMN  first_name VARCHAR(60);
ALTER TABLE users ADD COLUMN  last_name VARCHAR(60);
ALTER TABLE users ADD CONSTRAINT  unique_username UNIQUE(username);
