ALTER TABLE tweetsimages DROP CONSTRAINT IF EXISTS tweetsimages_primary_key;
ALTER TABLE tweetsimages
ADD CONSTRAINT tweetsimages_primary_key PRIMARY KEY (id);
