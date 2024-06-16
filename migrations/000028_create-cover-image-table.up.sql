CREATE TABLE IF NOT EXISTS cover_images(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    image BYTEA,
    user_id UUID REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
);
