ALTER TABLE tweetsimages
ADD COLUMN IF NOT EXISTS tweetId UUID REFERENCES tweets(id);