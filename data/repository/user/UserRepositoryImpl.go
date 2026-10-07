package user

import (
	"context"
	"database/sql"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"
	userRepo "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/user"
)

type repository struct {
	db      *sql.DB
	context context.Context
}

func NewUserRepository(db *sql.DB) userRepo.Repository {
	return &repository{
		db:      db,
		context: context.Background(),
	}
}

func (i repository) AddUser(value entities.SignUpForm) (*entities.UserID, error) {
	var dest string

	query := `
		INSERT INTO users(
			first_name,
			username,
			passwo
			($1, $2, $3) RETURNING id
	`

	if err := i.db.QueryRowContext(
		i.context,
		query,
		value.FirstName,
		value.Username,
		value.Password,
	).Scan(&dest); err != nil {
		return nil, err
	}

	return entities.ToUserID(dest)

}

func (i repository) FindUserCredentialsByUsername(username string) (*entities.UserCredential, error) {
	var dest *entities.UserCredential

	query := `
		SELECT
			id,
			password
		FROM users WHERE username=$1
	`

	if err := i.db.QueryRowContext(
		i.context,
		query,
		username,
	).Scan(&dest); err != nil {
		return nil, err
	}

	return dest, nil
}
