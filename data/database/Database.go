package database

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
)

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
