CREATE TABLE users(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	first_name TEXT NOT NULL,
	username TEXT UNIQUE NOT NULL,
	password_hash BYTEA NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE sessions(
	token_hash BYTEA PRIMARY KEY,
	user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	device_name TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE articles(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	title VARCHAR(200) NOT NULL,
	content TEXT NOT NULL,
	is_hide BOOLEAN NOT NULL DEFAULT false,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);