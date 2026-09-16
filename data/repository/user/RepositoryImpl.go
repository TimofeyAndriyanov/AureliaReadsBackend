package user

import (
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository"
)

type repositoryImpl struct{}

func NewUserRepository() repository.UserRepository {
	return &repositoryImpl{}
}

func (i repositoryImpl) AddUser(value entities.SignUpForm) entities.UserID {
	panic("")
}

func (i repositoryImpl) FindUserCredentialsByUsername(username string) (*entities.UserCredential, bool) {
	panic("")
}
