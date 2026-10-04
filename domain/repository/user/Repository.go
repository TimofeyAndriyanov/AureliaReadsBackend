package user

import "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"

type Repository interface {
	AddUser(value entities.SignUpForm) (*entities.UserID, error)
	FindUserCredentialsByUsername(username string) (*entities.UserCredential, error)
}
