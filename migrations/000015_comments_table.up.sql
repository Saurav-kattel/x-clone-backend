CREATE TABLE IF NOT EXISTS comments (
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
comment TEXT,
user_id UUID REFERENCES users(id),
tweet_id UUID REFERENCES tweets(id),
parent_comment_id UUID REFERENCES comments(id)
);
