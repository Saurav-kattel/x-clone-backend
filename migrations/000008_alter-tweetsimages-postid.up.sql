ALTER TABLE tweetsimages
ADD COLUMN tweetId UUID REFERENCES tweets(id);