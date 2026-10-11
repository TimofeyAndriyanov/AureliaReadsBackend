package user

import (
	"context"
	"database/sql"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"
	userRepo "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/user"
)

type sessionRepository struct {
	db      *sql.DB
	context context.Context
}

func NewSessionRepository(db *sql.DB) userRepo.SessionRepository {
	return &sessionRepository{
		db:      db,
		context: context.Background(),
	}
}

func (i sessionRepository) Add(uid entities.UserID, deviceName string, tokenHash entities.TokenHash) error {
	_ = `
		INSERT INTO sessions (
			token_hash,
			user_id,
			device_name
		) VALUES($1, $2, $3)
	`
	panic("")
}

func (i sessionRepository) DoesSessionExist(token entities.TokenHash) (bool, error) {
	_ = ""
	panic("")
}

func (i sessionRepository) GetSessionsByUserID(id entities.UserID) ([]entities.SessionItem, error) {
	panic("")
}

func (i sessionRepository) Delete(token entities.TokenHash) error {
	query := "DELETE FROM sessions WHERE token_hash=$1"

	if err := i.db.QueryRowContext(
		i.context,
		query,
		token,
	).Err(); err != nil {
		return err
	}

	return nil
}
