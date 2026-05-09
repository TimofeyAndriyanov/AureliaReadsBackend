package repository

import (
	"AureliaReadsBackend/data/models"
	"AureliaReadsBackend/data/utils"
	"AureliaReadsBackend/domain/entities"
	"AureliaReadsBackend/domain/repository"
	domainUtils "AureliaReadsBackend/domain/utils"
)

type userRepositoryImpl struct{}

func NewUserRepository() repository.UserRepository {
	return &userRepositoryImpl{}
}

var users []models.User

func (i userRepositoryImpl) AddUser(value entities.SignUpForm) entities.UserID {
	id := domainUtils.NewUserID()

	newUser := models.User{
		Id:        id,
		FirstName: value.FirstName,
		Username:  value.Username,
		Password:  value.Password,
	}

	users = append(users, newUser)

	return id
}

func (i userRepositoryImpl) FindUserCredentialsByUsername(username string) (*entities.UserCredential, bool) {
	user, ok := utils.Find(users, func(u models.User) bool {
		return u.Username == username
	})

	if user == nil {
		return nil, false
	}

	return user.ToUserCredential(), ok
}
