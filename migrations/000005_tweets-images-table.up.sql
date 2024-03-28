CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS tweetsImages(
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    image BYTEA
);