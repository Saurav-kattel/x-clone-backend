CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS tweets(
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    content text,
    userId UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_tweets_userId ON tweets(userId);
CREATE INDEX IF NOT EXISTS idx_tweets_createdAt ON tweets(created_at);