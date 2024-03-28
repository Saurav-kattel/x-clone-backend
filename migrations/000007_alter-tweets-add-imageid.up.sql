ALTER TABLE tweets
ADD COLUMN imageId UUID REFERENCES tweetsImages(id);