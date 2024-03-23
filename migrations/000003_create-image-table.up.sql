CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS images(
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    userId UUID REFERENCES users(id),
    image BYTEA
)