
CREATE TABLE IF NOT EXISTS reply (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  reply TEXT,
  tweet_id UUID REFERENCES tweets(id)  ON DELETE CASCADE ON UPDATE CASCADE,
  replied_to UUID REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
  replied_from UUID REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
  parent_id UUID REFERENCES reply(id) ON DELETE CASCADE ON UPDATE CASCADE
);

