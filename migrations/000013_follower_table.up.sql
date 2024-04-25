CREATE TABLE IF NOT EXISTS followers (
  follower_id UUID REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
  followee_id UUID REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
);

