package repository

import "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"

type UserRepository interface {
	AddUser(value entities.SignUpForm) entities.UserID
	FindUserCredentialsByUsername(username string) (*entities.UserCredential, bool)
}
