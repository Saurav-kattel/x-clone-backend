-- Drop the existing foreign key constraint (if it exists)
ALTER TABLE tweetsimages DROP CONSTRAINT IF EXISTS tweetsimages_tweetid_fkey;
-- Re-add the foreign key constraint with ON DELETE CASCADE
ALTER TABLE tweetsimages
ADD CONSTRAINT tweetsimages_tweetid_fkey FOREIGN KEY (tweetid) REFERENCES tweets(id) ON DELETE CASCADE ON UPDATE CASCADE;