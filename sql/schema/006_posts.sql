-- +goose Up
CREATE TABLE posts (
	id UUID PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT NOT NULL,
	url TEXT NOT NULL UNIQUE,
	published TIMESTAMP NOT NULL,
	feed_id UUID REFERENCES feeds(id) ON DELETE CASCADE,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE posts;