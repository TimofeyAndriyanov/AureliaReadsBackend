package database

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var _ = `
	CREATE TABLE users(
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		first_name TEXT NOT NULL,
		username TEXT UNIQUE NOT NULL,
		password_hash BYTEA NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
	);

	CREATE TABLE sessions(
		token_hash BYTEA PRIMARY KEY,
		user_id UUID NOT NULL REFERNCES users(id) ON DELETE CASCADE,
		device_name TEXT NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
	);

	CREATE TABLE articles(
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		user_id UUID NOT NULL REFERNCES users(id) ON DELETE CASCADE,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		is_hide BOOLEAN NOT NULL DEFAULT false,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
	);
`

func NewDatabase(host, port, user, password, dbName string, logger *slog.Logger) (*sql.DB, error) {
	var connectString = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbName,
	)
	db, openErr := sql.Open("pgx", connectString)

	if openErr != nil {
		logger.Error("err", openErr)
		return nil, openErr
	}

	if pingErr := db.Ping(); pingErr != nil {
		logger.Error("err", pingErr)
		return nil, pingErr
	}

	return db, nil
}
